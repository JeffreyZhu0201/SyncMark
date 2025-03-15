/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-14 22:20:14
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-15 19:14:35
 * @FilePath: \Smart-Snap-AI\Go-backend\main.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package main

import (
	"Go-backend/handlers"
	"Go-backend/middleware"
	"Go-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	utils.InitDB()

	// 设置路由
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	info := r.Group("/info")
	{
		info.GET("/", middleware.JWTAuthMiddleware, handlers.GetUserInfo)
	}

	r.GET("/", middleware.JWTAuthMiddleware, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动服务器
	if err := r.Run(); err != nil {
		panic(err)
	}
}
