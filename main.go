package main

import (
	"RcChat/mapper"
	"RcChat/models"
	"RcChat/router"
)

func main() {
	err := mapper.InitMysql()
	if err != nil {
		return
	}
	defer mapper.Open.Close()
	mapper.Open.AutoMigrate(&models.UserBasic{})
	r := router.Router()
	r.Run("localhost:8080")
}
