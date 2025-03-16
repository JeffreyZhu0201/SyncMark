package handlers

import (
	"Go-backend/models"
	"Go-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API_DeepSeek(question string, character string) (string, error) {
	params := map[string]any{
		"content":   question,
		"character": character,
	}

	res, err := utils.FetchPost("https://api.qster.top/API/v2/DeepSeek", params)

	if err != nil {
		return "", err
	}

	return res, nil

}

func HandleDeepSeek(c *gin.Context) {
	question := c.PostForm("question")
	character := c.PostForm("character")

	// 4. 调用 DeepSeek API
	deepSeekResult, err := API_DeepSeek(question, character)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "DeepSeek error"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 200, Message: "success", Data: deepSeekResult})
}
