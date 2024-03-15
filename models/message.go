package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FormId   string //发送者
	TargetId string //接收者
	Type     string //消息类型 群聊私聊广播
	Media    int    //消息类型 文字图片音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}
