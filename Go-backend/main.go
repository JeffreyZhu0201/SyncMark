/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-14 22:20:14
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-14 22:20:24
 * @FilePath: \Smart-Snap-AI\Go-backend\main.go
 * @Description: File Description Here...
 * 
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved. 
 */
package main

import "github.com/gin-gonic/gin"

func main() {
    // 初始化 Gin 引擎
    r := gin.Default()

    // 定义路由：GET 请求
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, Gin!",
        })
    })

    // 启动服务，默认端口 8080
    r.Run() // 或指定端口 r.Run(":3000")
}