/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-04-17 22:34:07
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-04-17 22:34:13
 * @FilePath: \SyncMark\Go-backend\handlers\Annotation.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package handlers

import (
	"Go-backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAnnotation(c *gin.Context) {
	var ann models.Annotation
	if err := c.ShouldBindJSON(&ann); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ann.CreatedAt = time.Now().Unix()
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&ann).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusOK, ann)
}

func ListAnnotations(c *gin.Context) {
	pageURL := c.Query("page_url")
	db := c.MustGet("db").(*gorm.DB)
	var anns []models.Annotation
	if err := db.Where("page_url = ?", pageURL).Find(&anns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	c.JSON(http.StatusOK, anns)
}

func DeleteAnnotation(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Delete(&models.Annotation{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": id})
}
