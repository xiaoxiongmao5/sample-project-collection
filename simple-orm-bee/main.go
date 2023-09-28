/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-28 13:55:09
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-28 23:52:31
 * @FilePath: /simple-orm-bee/main.go
 */
package main

import (
	"fmt"
	_ "simple-orm-bee/routers"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/client/orm/filter/bean"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

// User -
type User struct {
	//orm 为字段设置 DB 列的名称
	ID           int       `orm:"column(id);auto"`            //auto显示指定一个字段为自增主键，该字段必须是 int, int32, int64, uint, uint32, 或者 uint64
	UserAccount  string    `orm:"column(userAccount);unique"` //为单个字段增加 unique键
	AccessKey    string    `orm:"column(accessKey);index"`    //为单个字段增加索引
	SecretKey    string    `orm:"column(secretKey)"`
	UserRole     string    `orm:"column(userRole);default(admin)"`
	UserName     string    `orm:"column(userName);null"` //null代表ALLOW NULL
	UserAvatar   string    `orm:"column(userAvatar);null"`
	Gender       int8      `orm:"column(gender);null"`
	UserPassword string    `orm:"column(userPassword)"`
	CreateTime   time.Time `orm:"column(createTime);auto_now_add;type(datetime)"` //auto_now_add 第一次保存时才设置时间
	UpdateTime   time.Time `orm:"column(updateTime);auto_now;type(datetime)"`     //auto_now 每次model保存时都会对时间自动更新
	IsDelete     int32     `orm:"column(isDelete)"`
	Test         string    `orm:"-"` //- 即可忽略模型中的字段
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}

// 自定义表名
func (u *User) TableName() string {
	return "user"
}

// 多字段索引(联合索引)
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"ID", "UserName"},
	}
}

// 多字段唯一键
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"AccessKey", "SecretKey"},
	}
}

var O orm.Ormer

func init() {
	// 注册 ORM 模型。new(User) 创建了一个新的 User 模型对象并注册到 ORM 中。这样 ORM 知道了要操作哪个数据表和数据表的结构。
	orm.RegisterModel(new(User))

	// 注册默认的数据库连接。告诉 ORM 使用 MySQL 数据库，数据库连接字符串为 "root:xxx"，并且指定了连接别名为 "default"。这个连接别名会在后续的数据库操作中使用。
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/xapi?charset=utf8")

	// 创建一个 ORM 对象 o，用于执行数据库操作。
	O = orm.NewOrm()
}
func main() {
	// 执行数据库表结构同步操作。告诉 ORM 在默认数据库连接上执行同步操作，第二个参数 false 表示不强制删除已存在的表，第三个参数 true 表示打印同步操作的日志。这通常在应用程序启动时执行，以确保数据库表结构与 ORM 模型定义一致。
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("数据库表结构同步操作失败[orm.RunSyncdb] err=", err.Error())
	}

	// 创建一个默认值过滤器链构建器。默认值过滤器用于在插入数据时为字段设置默认值。这里创建了一个新的构建器，第二个参数 true 表示启用默认值过滤器，第三个参数 true 表示打印过滤器日志。
	builder := bean.NewDefaultValueFilterChainBuilder(nil, true, true)
	// 将默认值过滤器链添加到全局过滤器链中，以便在插入数据时应用默认值过滤器。
	orm.AddGlobalFilterChain(builder.FilterChain)

	id := Add()

	userinfo, err := GetById(id)
	if err == nil {
		fmt.Println("userinfo=", userinfo)
	}

	beego.Run()
}

func GetById(id int64) (*User, error) {
	qs := O.QueryTable(new(User))

	var userInfo User
	err := qs.Filter("id", id).Filter("isDelete", 0).One(&userInfo) //注意这里不能是 var userInfo *User，否则报错
	if err == orm.ErrMultiRows {
		fmt.Printf("user 表中存在 id=[%d] 的多条记录, qs.One err=[%v] \n", id, err.Error())
		return nil, err
	}
	if err == orm.ErrNoRows {
		fmt.Printf("user 表没有找到 id=[%d] 的记录, qs.One err=[%v] \n", id, err.Error())
		return nil, err
	}
	return &userInfo, nil
}

func ListByIds(ids []int64) ([]*User, error) {
	qs := O.QueryTable(new(User))

	var users []*User
	qs = qs.Filter("id__in", ids).Filter("isDelete", 0)
	_, err := qs.All(&users) //这里可以是 var users []*User 或者 var users []User
	if err != nil {
		fmt.Printf("User ListByIds qs.All error: %v \n", err.Error())
		return users, err
	}
	return users, nil
}

func Add() int64 {
	// 创建了一个新的 User 模型对象，准备插入到数据库中。
	user := new(User)
	user.UserAccount = "xiaoxiong"

	// 执行数据库插入操作，将 user 对象插入到数据库中。num 变量表示插入的记录数，err 变量表示操作中的任何错误。
	num, err := O.Insert(user)
	if err != nil {
		fmt.Println("数据库插入操作失败[o.Insert(user)] err=", err.Error())
	}
	fmt.Println("数据库插入操作成功[o.Insert(user) succeed] num=", num)
	return num
}
