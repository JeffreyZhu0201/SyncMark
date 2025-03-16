// DeepSeek.go 处理与DeepSeek API相关的请求
// 包含两个主要函数:
// 1. API_DeepSeek: 调用DeepSeek API获取回答
// 2. HandleDeepSeek: 处理前端DeepSeek请求的HTTP处理器

// API_DeepSeek 通过DeepSeek API获取对问题的回答
// 参数:
//   - question: 用户提出的问题字符串
// 返回:
//   - string: API的响应内容
//   - error: 如果请求过程中出现错误则返回错误信息,否则返回nil

// HandleDeepSeek 处理来自前端的DeepSeek请求
// 通过gin.Context接收POST请求中的question参数
// 调用API_DeepSeek获取结果
// 将结果解析为JSON并返回给客户端
// 如果处理过程中出现错误,返回相应的错误状态码和信息
package handlers

import (
	"Go-backend/models"
	"Go-backend/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API_DeepSeek(question string) (string, error) {
	params := map[string]any{
		"content": question,
	}

	res, err := utils.FetchGet("https://api.qster.top/API/v2/DeepSeek", params)

	if err != nil {
		return "", err
	}

	return res, nil

}

func HandleDeepSeek(c *gin.Context) {
	question := c.PostForm("question")

	// 4. 调用 DeepSeek API
	deepSeekResult, err := API_DeepSeek(question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "DeepSeek API error"})
		return
	}

	var parsedResult interface{}
	err = json.Unmarshal([]byte(deepSeekResult), &parsedResult)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "DeepSeek error"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 200, Message: "success", Data: parsedResult})
}
