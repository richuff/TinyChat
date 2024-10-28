package utils

import (
	"RcChat/mapper"
	"RcChat/models"
	"fmt"
	"github.com/spf13/viper"
)

func InitMysql() {

	err := mapper.InitMysql(viper.GetString("mysql.dns"))
	if err != nil {
		return
	}
	mapper.Open.AutoMigrate(&models.UserBasic{})
	user := models.UserBasic{}
	mapper.Open.Find(&user)
	fmt.Println(user)
}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	fmt.Println(viper.Get("mysql"))
}
