/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-17 22:24:39
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-04-17 22:34:03
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

type Annotation struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	User      string `json:"user"`
	PageURL   string `json:"page_url"`
	Text      string `json:"text"`
	Quote     string `json:"quote"`
	Range     string `json:"range"` // 可存储选中文本的位置信息（如XPath或offset）
	CreatedAt int64  `json:"created_at"`
}
