package handlers

import (
    "Go-backend/models"
    "Go-backend/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

// CreateRoom handles the creation of a new room
func CreateRoom(c *gin.Context) {
    var room models.Room
    if err := c.ShouldBindJSON(&room); err != nil {
        c.JSON(http.StatusBadRequest, models.Response{Code: 400, Message: "Invalid request"})
        return
    }

    if err := utils.DB.Create(&room).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "Error creating room"})
        return
    }

    c.JSON(http.StatusOK, models.Response{Code: 200, Message: "Room created successfully"})
}

// DeleteRoom handles the deletion of a room
func DeleteRoom(c *gin.Context) {
    id := c.Param("id")
    if err := utils.DB.Delete(&models.Room{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "Error deleting room"})
        return
    }

    c.JSON(http.StatusOK, models.Response{Code: 200, Message: "Room deleted successfully"})
}

// GetRooms handles fetching all rooms
func GetRooms(c *gin.Context) {
    var rooms []models.Room
    if err := utils.DB.Find(&rooms).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "Error fetching rooms"})
        return
    }

    c.JSON(http.StatusOK, models.Response{Code: 200, Message: "Success", Data: rooms})
}