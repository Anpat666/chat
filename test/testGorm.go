package main

import (
	"chat/models"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:135423@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})

	// Create
	var user = models.UserBasic{
		Name:          "test",
		HashPassword:  "123456",
		Phone:         "8888888",
		LoginTime:     time.Now().Format("2006-01-02 15:04:05"),
		LogOutTime:    time.Now().Format("2006-01-02 15:04:05"),
		HeartbeatTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	db.Create(&user)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	data := rdb.Get(context.Background(), "anpat").Val()
	fmt.Println(data)
}
