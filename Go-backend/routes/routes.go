/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-17 23:08:33
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-17 23:19:47
 * @FilePath: \SyncMark\Go-backend\routes\routes.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package routes

import (
	"Go-backend/handlers"
	"Go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	//认证相关路由
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// ai相关路由
	aiInterface := r.Group("/ai")
	{
		aiInterface.POST("/ocr", middleware.JWTAuthMiddleware, handlers.HandleUploadImg)
		aiInterface.POST("/deepseek", middleware.JWTAuthMiddleware, handlers.HandleDeepSeek)
	}

	// 测试路由
	r.GET("/", middleware.JWTAuthMiddleware, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// WebSocket 路由
	r.GET("/ws", handlers.HandleWebSocket)
}
