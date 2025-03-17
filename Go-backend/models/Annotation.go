/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-17 22:24:39
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-17 22:24:51
 * @FilePath: \SyncMark\Go-backend\models\Annotation.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */

package models

type Annotation struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	URLHash   string `json:"url_hash"`
	Content   string `json:"content"`
	Position  string `json:"position"`
	UserID    uint   `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
}
