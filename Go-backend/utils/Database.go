package utils

import (
	"Go-backend/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接并返回数据库实例
func InitDB() *gorm.DB {
	dsn := "root:qQ13527503399@tcp(127.0.0.1:3306)/smart-snap-ai?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移数据库表
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	return DB
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
