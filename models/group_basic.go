package models

import "gorm.io/gorm"

// GroupBasic  群表
type GroupBasic struct {
	gorm.Model
	Name   string
	OwerId uint
	Icon   string
	Type   int
	Desc   string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
