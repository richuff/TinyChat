package models

import (
	"RcChat/mapper"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    /*用户名*/
	Password      string    /*密码*/
	Phone         string    /*电话号码*/
	Email         string    /*邮箱*/
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

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 3)
	mapper.Open.Find(&data)
	for i, v := range data {
		fmt.Println(i, v)
	}
	return data
}
