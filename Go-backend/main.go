/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-14 22:20:14
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-17 22:23:22
 * @FilePath: \SyncMark\Go-backend\main.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package main

import (
	"Go-backend/routes"
	"Go-backend/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("加载环境变量失败: %v", err)
	}

	// 初始化数据库
	utils.InitDB()

	// 初始化路由
	r := gin.Default()
	routes.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(); err != nil {
		panic(err)
	}
}
