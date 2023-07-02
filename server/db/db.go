package db

import (
	"context"
	"fmt"
	"server/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MYSQL *gorm.DB
	Redis *redis.Client
)

func init() {
	//
	InitMyql()
	//
	InitRedis()
}

func InitMyql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/ai_box?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.MySql.Username,
		config.Cfg.MySql.Password,
		config.Cfg.MySql.Host,
		config.Cfg.MySql.Port,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.Cfg.MySql.Level)),
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	MYSQL = db
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DB,
	})

	result := Redis.Ping(context.Background())
	if result.Val() != "PONG" {
		panic("Redis Connect Fail")
	}
}
