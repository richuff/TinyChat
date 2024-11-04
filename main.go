package main

import (
	"RcChat/router"
	"RcChat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	//mapper.Open.AutoMigrate(&models.UserBasic{})
	utils.InitRedis()

	r := router.Router()
	r.Run("localhost:8080")
}
