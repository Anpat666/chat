package main

import (
	"fmt"
	"lottery/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:135423@tcp(localhost:3306)/lottery?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接MySQL失败:", err)
	}

	DB.AutoMigrate(&models.Player{})
}
