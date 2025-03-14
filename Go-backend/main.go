/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-14 22:20:14
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-14 23:59:32
 * @FilePath: \Smart-Snap-AI\Go-backend\main.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package main

import (
	"Go-backend/handlers"
	"Go-backend/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "root:qQ13527503399@tcp(127.0.0.1:3306)/smart-snap-ai?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{})
}

func main() {
	initDB()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	r.Run()
}
