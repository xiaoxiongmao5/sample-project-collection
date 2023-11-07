package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
Gorm 读写数据库
* Orm 就是把一个对象跟一个表结构进行关联、映射
* Gorm 是Go中用的最多的 MySQL 驱动器。
* Gorm 默认的结构体名称的“蛇形复数”等于表名。
* 	驼峰：User、UserName
* 	蛇形：user、user_name
*   蛇形复数：users
*/

// 默认 User 这个结构体对应的表名是 users
type User struct {
	Id      int    //id 主键 默认情况下，Gorm 认为 Id 就是主键。如果 Id 不是主键的话，比如需要`gorm:"column:aid,primaryKey"`
	Uid     int    //uid
	KeyWord string `gorm:"column:keywords"` //转为蛇形是 keyword，如果和表中字段不同，需要显式指定对应的字段名
	City    string //city
}

// 显式指定表名
func (User) TableName() string {
	return "user"
}

// 读
func read(client gorm.DB, city string) *User {
	// var users []User
	// 使用 ？会更加安全，防止 SQL注入攻击
	// client.Select("id,city").Where("city=?", city).Find(&users)
	// if len(user) > 0 {
	// 	return &users[0]
	// } else {
	// 	return nil
	// }

	var user User
	// err := client.Select("id,city").Where("city=?", city).First(&user).Error //First 不管前面是几个，一个或多个，只取第一个

	user.Id = 1
	err := client.Select("id,city").Where("city=?", city).First(&user).Error //会隐含一个where条件 id=1&city=北京

	// err = client.Select("id,city").Where("city=?", city).Limit(1).Take(&user).Error   //Take 在前面返回多条的基础上，随便取一个
	checkErr(err)
	return &user
}

// 插入
func insert(client gorm.DB) {
	user := User{
		Id:      3,
		KeyWord: "golang",
		City:    "上海",
	}
	err := client.Create(user).Error
	checkErr(err)
}

// 批量插入
func insertList(client gorm.DB) {
	users := make([]User, 0)
	for i := 4; i < 10; i++ {
		users = append(users, User{
			Id:      i,
			Uid:     i,
			KeyWord: "golang",
			City:    "北京",
		})
	}

	err := client.CreateInBatches(users, len(users)).Error
	checkErr(err)
}

// 更新
func update(client gorm.DB) {
	err := client.Model(User{}).Where("id=?", 3).Update("city", "贵阳").Error
	checkErr(err)

	err = client.Model(User{}).Where("id=?", 4).Updates(map[string]interface{}{
		"city":   "湖南",
		"gender": "男",
	}).Error
	checkErr(err)
}

// 删除
func delete(client gorm.DB) {
	err := client.Where("id=?", 7).Delete(User{}).Error
	checkErr(err)
}

func main() {
	dataSourceName := "root:@tcp(127.0.0.1:3306)/test_gorm?charset=utf8&parseTime=true"
	// 建立链接
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkErr(err)
	user := read(*client, "北京")
	if user != nil {
		fmt.Printf("%+v\n", user)
	} else {
		fmt.Println("无结果")
	}

	// insert(*client)
	// insertList(*client)

	// update(*client)

	delete(*client)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
