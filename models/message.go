package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromID   string //发送者
	TargetId string //接收者ID
	Type     string //消息类型
	Media    int    //消息类型
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其它数组统计
}

func (table *Message) TableName() string {
	return "Message_basic"
}
