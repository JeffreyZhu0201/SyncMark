package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 传入目标url和参数map，返回post请求的结果
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

// 传入目标url和参数map，返回get请求的结果
func FetchGet(targetUrl string, params map[string]any) (string, error) {
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
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
