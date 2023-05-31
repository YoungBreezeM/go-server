package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MYSQL *gorm.DB

func init() {
	// 创建数据库连接
	dsn := "user:password@tcp(host:port)/database"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic("Failed to connect to database")
		fmt.Println(err)
	}

	// 设置数据库连接池参数等
	// db.DB().SetMaxOpenConns(100)
	// db.DB().SetMaxIdleConns(10)
	// ...

	MYSQL = db
}
