package db

import (
	"fmt"
	"server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MYSQL *gorm.DB

func init() {
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/ai_box?charset=utf8mb4&parseTime=True&loc=Local", config.Cfg.MySql.Username, config.Cfg.MySql.Password, config.Cfg.MySql.Host, config.Cfg.MySql.Port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.Cfg.MySql.Level)),
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	MYSQL = db
}
