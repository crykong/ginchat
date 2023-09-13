package models

import "gorm.io/gorm"

// Contact 关系表
type Contact struct {
	gorm.Model
	OwerId   int64 //谁的关系
	TargetId uint  //对应的谁
	Type     int   //对应类型  0 1 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact_basic"
}
