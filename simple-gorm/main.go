package main

import "xj/simple-gorm/db"

/**
Gorm 读写数据库
* Orm 就是把一个对象跟一个表结构进行关联、映射
* Gorm 是Go中用的最多的 MySQL 驱动器。
* Gorm 默认的结构体名称的“蛇形复数”等于表名。
* 	驼峰：User、UserName
* 	蛇形：user、user_name
*   蛇形复数：users
*/

func main() {
	// db.DoAutoMigrate()
	// db.DoQuickCurd()
	db.DoZero()
	// dataSourceName := "root:@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// // loc 设置本地系统的时间，前提 parseTime=True
	// // 更多参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// // 建立链接
	// client, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	// // client, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{Logger:logger.Default.LogMode(logger.Info)})  打印sql日志
	// db.CheckErr(err)
	// user := db.Read(*client, "北京")
	// if user != nil {
	// 	fmt.Printf("%+v\n", user)
	// } else {
	// 	fmt.Println("无结果")
	// }

	// db.Insert(*client)
	// db.InsertList(*client)
	// db.Update(*client)
	// db.Delete(*client)
}
