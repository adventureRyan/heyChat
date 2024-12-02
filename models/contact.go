package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerId  uint
	TargetId uint
	Type     int
	Desc     string // 预留字段
}

func (table *Contact) TableName() string {
	return "contact"
}
