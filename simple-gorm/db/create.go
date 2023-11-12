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

type User2 struct {
	gorm.Model
	Name     string
	Age      uint8
	Birthday time.Time
}

func GetDb() *gorm.DB {
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
		panic("failed to connect database")
	}
	db.AutoMigrate(&User2{})
	return db
}

// 通过数据的指针来创建
func DoCreate(db *gorm.DB) {
	user := User2{Name: "xiaoa", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)       //通过数据的指针来创建
	fmt.Println(user.ID)             //返回插入数据的主键
	fmt.Println(result.Error)        //返回 error
	fmt.Println(result.RowsAffected) //返回插入记录的条数
}

// 用指定的字段创建记录
func DoCreateWithSelect(db *gorm.DB) {
	user := User2{Name: "xiaob", Age: 19, Birthday: time.Now()}
	db.Select("Name", "Age", "CreateAt").Create(&user)
	// INSERT INTO `user2` (`created_at`,`updated_at`,`name`,`age`) VALUES ('2023-11-12 23:02:34.929','2023-11-12 23:02:34.929','xiaob',19)
}

// 创建记录并忽略指定字段值（指定的字段被忽略掉，不会赋值创建）
func DoCreateWithOmit(db *gorm.DB) {
	user := User2{Name: "xiaoc", Age: 20, Birthday: time.Now()}
	db.Omit("Name", "Age", "CreateAt").Create(&user)
	//  INSERT INTO `user2` (`created_at`,`updated_at`,`deleted_at`,`birthday`) VALUES ('2023-11-12 23:06:41.381','2023-11-12 23:06:41.381',NULL,'2023-11-12 23:06:41.381')
}

// 批量插入
func DoManyCreate(db *gorm.DB) {
	// time.Parse 函数用于将字符串解析为时间。
	// "2006-01-02" 是一个时间格式字符串，它定义了时间的布局，例如 "2006" 表示年，"01" 表示月，"02" 表示日。
	// "1998-07-12" 是要被解析的时间字符串。通过将这个字符串和时间布局一起传递给 time.Parse 函数，Go会尝试解析这个字符串，并返回相应的时间对象。
	// time.Parse("2006-01-02", "1998-07-12") 将返回一个时间对象，表示 "1998-07-12" 这个日期。
	t, _ := time.Parse("2006-01-02", "1998-07-12")
	var users = []User2{
		{Name: "xiaox", Birthday: t},
		{Name: "xiaoy", Birthday: t},
		{Name: "xiaoz", Birthday: t},
	}
	db.Create(&users) // create 不仅可以放对象指针，也能放切片的指针。批量创建
	for _, user := range users {
		fmt.Println(user.ID)
	}
	// 因为Mysql5.7版本及以上版本的datetime值不能为'0000-00-00 00:00:00',
	// //处理方法：
	// 修改mysql.ini
	// 在[mysqld]添加一项：
	// sql_mode=NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO
}
