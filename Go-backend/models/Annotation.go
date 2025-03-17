/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-17 22:24:39
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-18 00:08:24
 * @FilePath: \SyncMark\Go-backend\models\Annotation.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
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

import (
	"gorm.io/gorm"
)

type Annotation struct {
	gorm.Model
	RoomID  uint   `json:"room_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
	PosX    int    `json:"pos_x"`
	PosY    int    `json:"pos_y"`
}
