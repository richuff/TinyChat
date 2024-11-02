package models

import (
	"RcChat/mapper"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    /*用户名*/
	Password      string    /*密码*/
	Phone         string    `valid:"matches(^1[3-9]{1}\\d{9}$)"` /*电话号码*/
	Email         string    `valid:"email"`                      /*邮箱*/
	Identify      string    /*验证*/
	ClientIp      string    /*客户端ip*/
	ClientPort    string    /*客户端端口*/
	LoginTime     time.Time /*登录时间*/
	HeartBeatTime time.Time /*心跳时间*/
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"` /*退出登录时间*/
	IsLoginOut    bool      /*是否退出*/
	DeviceInfo    string    /*设备信息*/
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func FindUserByPhone(data *UserBasic) *gorm.DB {
	return mapper.Open.Create(&data)
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 3)
	mapper.Open.Find(&data)
	for i, v := range data {
		fmt.Println(i, v)
	}
	return data
}

func CreateUser(data *UserBasic) *gorm.DB {
	return mapper.Open.Create(&data)
}

func DeleteUser(data *UserBasic) *gorm.DB {
	return mapper.Open.Delete(&data)
}

func UpdateUser(data UserBasic) *gorm.DB {
	return mapper.Open.Model(&UserBasic{}).Where("id = ?", data.ID).Updates(UserBasic{Name: data.Name, Password: data.Password, Email: data.Email, Phone: data.Phone})
}
