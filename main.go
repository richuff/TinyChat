package main

import (
	"RcChat/router"
	"RcChat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()
	r.Run("localhost:8080")
}
