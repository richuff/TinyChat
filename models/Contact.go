package models

import (
	"RcChat/mapper"
	"gorm.io/gorm"
	"log"
	"strconv"
)

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

func SearchFriend(userId uint64) []UserBasic {
	contacts := make([]Contact, 0)
	mapper.Open.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	objIds := make([]uint64, 0)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
		log.Println(v)
	}
	users := make([]UserBasic, 0)
	mapper.Open.Where("id in ?", objIds).Find(&users)
	return users
}

func AddFriend(userId uint, targetId uint) (int, string) {
	if targetId == userId {
		return -1, "不能添加自己"
	}
	user := UserBasic{}
	if targetId != 0 {
		user = FindUserById(targetId)
		if user.Identify != "" {
			contact0 := Contact{}
			mapper.Open.Where("owner_id = ? and target_id = ? and type = 1", userId, targetId).Find(&contact0)
			if contact0.TargetId != 0 {
				return -1, "不能重复添加"
			}

			//创建事务
			tx := mapper.Open.Begin()
			//事务异常都会Rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			contact := Contact{
				OwnerID:  userId,
				TargetId: targetId,
				Type:     1,
			}

			if err := mapper.Open.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}

			contact = Contact{
				OwnerID:  targetId,
				TargetId: userId,
				Type:     1,
			}

			if err := mapper.Open.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			//提交事务
			tx.Commit()
			return 0, "添加好友成功"
		}
		return -1, "找不到该用户"
	}
	return -1, "输入的格式错误"
}

func JoinCommunity(u uint, id string) (int, string) {
	contact := Contact{}
	contact.OwnerID = u
	contact.Type = 2
	community := Community{}
	mapper.Open.Where("id = ? or name = ?", id, id).Find(&community)
	if community.Name == "" {
		return -1, "该群不存在"
	}
	mapper.Open.Where("owner_id = ? and target_id = ? and type = 2", u, id).Find(&contact)
	if contact.ID != 0 {
		return -1, "已经添加过该群"
	} else {
		ComId, _ := strconv.Atoi(id)
		contact.TargetId = uint(ComId)
		mapper.Open.Create(&contact)
		return 0, "创建成功"
	}
}

func FindUserIdsByTargetId(targetId uint) []int64 {
	contacts := make([]Contact, 20)
	res := make([]int64, 0)
	mapper.Open.Where("target_id = ? and type = 2", targetId).Find(&contacts)
	for _, v := range contacts {
		res = append(res, int64(v.ID))
	}
	return res
}
