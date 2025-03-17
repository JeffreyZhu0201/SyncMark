package handlers

import (
    "Go-backend/models"
    "Go-backend/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

// CreateAnnotation handles the creation of a new annotation
func CreateAnnotation(c *gin.Context) {
    var annotation models.Annotation
    if err := c.ShouldBindJSON(&annotation); err != nil {
        c.JSON(http.StatusBadRequest, models.Response{Code: 400, Message: "Invalid request"})
        return
    }

    if err := utils.DB.Create(&annotation).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "Error creating annotation"})
        return
    }

    c.JSON(http.StatusOK, models.Response{Code: 200, Message: "Annotation created successfully"})
}

// DeleteAnnotation handles the deletion of an annotation
func DeleteAnnotation(c *gin.Context) {
    id := c.Param("id")
    if err := utils.DB.Delete(&models.Annotation{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "Error deleting annotation"})
        return
    }

    c.JSON(http.StatusOK, models.Response{Code: 200, Message: "Annotation deleted successfully"})
}

// GetAnnotations handles fetching all annotations for a specific room
func GetAnnotations(c *gin.Context) {
    roomId := c.Param("roomId")
    var annotations []models.Annotation
    if err := utils.DB.Where("room_id = ?", roomId).Find(&annotations).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "Error fetching annotations"})
        return
    }

    c.JSON(http.StatusOK, models.Response{Code: 200, Message: "Success", Data: annotations})
}