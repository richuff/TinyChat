package main

import (
	"RcChat/router"
	"RcChat/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	//mapper.Open.AutoMigrate(&models.UserBasic{})
	//mapper.Open.AutoMigrate(&models.Message{})    //消息
	//mapper.Open.AutoMigrate(&models.GroupBasic{}) //群信息
	//mapper.Open.AutoMigrate(&models.Contact{})
	//mapper.Open.AutoMigrate(&models.Community{})

	r := router.Router()
	err := r.Run(viper.GetString("server.port"))
	if err != nil {
		return
	}
}
