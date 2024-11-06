package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID  uint //谁的关系消息
	TargetId uint //对应的谁
	Type     int  //对应的类型
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
