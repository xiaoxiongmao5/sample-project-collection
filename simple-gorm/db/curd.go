package db

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

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
func Read(client gorm.DB, city string) *User {
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
	CheckErr(err)
	return &user
}

// 插入
func Insert(client gorm.DB) {
	user := User{
		Id:      3,
		KeyWord: "golang",
		City:    "上海",
	}
	err := client.Create(user).Error
	CheckErr(err)
}

// 批量插入
func InsertList(client gorm.DB) {
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
	CheckErr(err)
}

// 更新
func Update(client gorm.DB) {
	err := client.Model(User{}).Where("id=?", 3).Update("city", "贵阳").Error
	CheckErr(err)

	err = client.Model(User{}).Where("id=?", 4).Updates(map[string]interface{}{
		"city":   "湖南",
		"gender": "男",
	}).Error
	CheckErr(err)
}

// 删除
func Delete(client gorm.DB) {
	err := client.Where("id=?", 7).Delete(User{}).Error
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
