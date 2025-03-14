/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-14 23:10:39
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-14 23:27:46
 * @FilePath: \Smart-Snap-AI\Go-backend\models\models.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
