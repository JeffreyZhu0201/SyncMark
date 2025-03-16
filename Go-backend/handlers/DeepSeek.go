package handlers

import (
	"Go-backend/utils"

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

}
