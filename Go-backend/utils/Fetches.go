package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func FetchPost(targetUrl string, params map[string]any) (string, error) {

	//将map转为json格式
	jsonData, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(targetUrl, "application/json", bytes.NewBuffer(jsonData)) //将json转为字符串
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

// func FetchGet(targetUrl string, params map[string]any) (string, error) {

// 	//将map转为json格式
// 	jsonData, err := json.Marshal(params)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := http.Get(targetUrl, bytes.NewBuffer(jsonData)) //将json转为字符串
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// Read response
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(body), nil

// }
