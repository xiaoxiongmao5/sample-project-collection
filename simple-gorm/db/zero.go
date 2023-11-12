package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 表模型修改为 sql.NullString
type Product2 struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

// 表模型修改为 *string
type Product3 struct {
	gorm.Model
	Code  *string
	Price uint
}

// 零值问题
func DoZero() {
	dsn := "root:@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // 如果数据库不存在会报错
	}

	var product Product
	db.First(&product)
	fmt.Printf("product= %+v \n", product)

	db.Model(&product).Updates(Product{Price: 200, Code: ""}) //仅更新非零值字段
	// UPDATE `products` SET `updated_at`='2023-11-12 21:09:59.879',`price`=200 WHERE `products`.`deleted_at` IS NULL AND `id` = 2
	fmt.Printf("product= %+v \n", product) // 可以看到Code字段不会更新，这是合理的，因为如果零值字段也更新，Product表中好多数据都会被更新为空

	UseSqlNullString(db)
	UsePointer(db)
}

// 使用 sql.NullString 更新
func UseSqlNullString(db *gorm.DB) {
	// 创建表 product2
	db.AutoMigrate(Product2{})
	db.Create(&Product2{Price: 150})
	db.Create(&Product2{Price: 250})
	var product2 Product2
	// 修改语句为
	db.Model(&product2).Where("id=?", 5).Updates(Product2{Price: 200, Code: sql.NullString{"", true}})
	fmt.Printf("product2= %+v \n", product2)
}

// 使用指针解决
func UsePointer(db *gorm.DB) {
	// 创建表 product3
	db.AutoMigrate(Product3{})
	db.Create(&Product3{Price: 150})
	db.Create(&Product3{Price: 250})
	var product3 Product3
	empty := ""
	// 修改语句为
	db.Model(&product3).Where("id=?", 5).Updates(Product3{Price: 200, Code: &empty})
	fmt.Printf("product3= %+v \n", product3)
}
