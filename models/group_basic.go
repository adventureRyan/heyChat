package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string // 预留字段
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
