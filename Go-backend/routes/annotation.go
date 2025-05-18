/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-04-17 22:34:16
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-04-17 22:34:26
 * @FilePath: \SyncMark\Go-backend\routes\annotation.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package routes

import (
	"Go-backend/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAnnotationRoutes(r *gin.Engine) {
	ann := r.Group("/api/annotation")
	{
		ann.POST("/", handlers.CreateAnnotation)
		ann.GET("/", handlers.ListAnnotations)
		ann.DELETE("/:id", handlers.DeleteAnnotation)
	}
}
