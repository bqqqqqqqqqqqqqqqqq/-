package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/wx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error:", err)
	}
	return db
}

func InitRedisDB() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
