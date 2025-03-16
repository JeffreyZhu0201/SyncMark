package handlers

import (
	"Go-backend/models"
	"Go-backend/utils"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API_OCR(Base64Image string) (string, error) {
	params := map[string]any{
		"token":       "www.tulingyun.com",
		"upfile_b64":  Base64Image,
		"return_text": 1,
	}

	res, err := utils.FetchPost("https://www.tulingyun.com/api/ocr", params)

	if err != nil {
		return "", err
	}

	return res, nil
}

func HandleUploadImg(c *gin.Context) {

	// Get the file from the form
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Message: "upload error"})
		return
	}
	// 2. 打开文件并读取内容
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 400, Message: "upload error"})
		return
	}
	defer openedFile.Close()

	fileBytes := make([]byte, file.Size)
	_, err = openedFile.Read(fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 400, Message: "upload error"})
		return
	}

	// 3. 转换为 Base64
	base64Str := base64.StdEncoding.EncodeToString(fileBytes)

	// 4. 调用 OCR API
	ocrResult, err := API_OCR(base64Str)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Message: "ocr error"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 200, Message: "success", Data: ocrResult})
}
