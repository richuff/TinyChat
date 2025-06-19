package main

import (
	"RcChat/router"
	"RcChat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	//mapper.Open.AutoMigrate(&models.UserBasic{})
	//mapper.Open.AutoMigrate(&models.Message{})    //消息
	//mapper.Open.AutoMigrate(&models.GroupBasic{}) //群信息
	//mapper.Open.AutoMigrate(&models.Contact{})
	utils.InitRedis()

	r := router.Router()
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
