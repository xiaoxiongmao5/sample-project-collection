package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DoQuickCurd() {
	// 日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别为info
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	dsn := "root:@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err) // 如果数据库不存在会报错
	}
	//1 迁移 表，可以写多个
	// db.AutoMigrate(&Product{})

	//2 插入数据
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "D43", Price: 150})

	//3 查询数据
	var product Product   // 定义空Product结构体对象
	db.First(&product, 1) // 根据整型主键查找
	fmt.Println("查询 id=1,", product)
	err = db.First(&product, "code = ?", "D43").Error // 查找 code 字段值为 D43 的记录
	if err != nil {
		fmt.Println("查询 code=D43 失败,", err)
	} else {
		fmt.Println("查询 code=D43,", product)
	}

	//4  更新 - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	fmt.Println("更新 price=200,", product)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F43"}) // 仅更新非零值字段
	fmt.Println("更新 price=200,code=F43,", product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "F42"})
	fmt.Println("更新 price=300,code=F42,", product)

	//5  Delete - 删除 product
	db.Delete(&product, 1) // 软删除
}
