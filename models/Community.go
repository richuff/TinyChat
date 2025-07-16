package models

import (
	"RcChat/mapper"
	"gorm.io/gorm"
	"log"
)

type Community struct {
	gorm.Model
	Name     string
	OwnerId  uint
	TargetId uint
	Image    string
	Desc     string
}

func (*Community) TableName() string {
	return "community"
}

func CreateCommunity(community Community) (int, string) {
	if community.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := mapper.Open.Create(&community).Error; err != nil {
		log.Println(err)
		return -1, "创建群聊失败"
	}
	return 0, "创建群聊成功"
}

func LoadCommunity(OwnerId uint) ([]Community, string) {
	data := make([]Community, 10)
	if err := mapper.Open.Where("owner_id=?", OwnerId).Find(&data).Error; err != nil {
		log.Println(err)
		return nil, "查询失败"
	}
	return data, "查询成功"
}
