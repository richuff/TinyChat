package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   string `json:"from_id"`   //发送者
	TargetId string `json:"target_id"` //接受者
	Type     string `json:"type"`      //消息来源类型 群聊 私聊 广播
	Media    int    `json:"media"`     //消息类型 文字 图片 音频 视频
	Content  string `json:"content"`
	Pic      string `json:"pic"`
	Desc     string `json:"desc"`
	Url      string `json:"url"`
	Amount   int    `json:"amount"` //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}
