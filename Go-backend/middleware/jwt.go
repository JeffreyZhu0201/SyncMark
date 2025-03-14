/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-14 23:11:18
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-14 23:56:50
 * @FilePath: \Smart-Snap-AI\Go-backend\config\jwt.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package middleware

import (
	"Go-backend/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) { // 修改此行
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
