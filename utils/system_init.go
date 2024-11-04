package utils

import (
	"RcChat/mapper"
	"fmt"
	"github.com/spf13/viper"
)

func InitRedis() {
	pong, err := mapper.InitRedis(viper.GetString("redis.addr"),
		viper.GetString("redis.password"), viper.GetInt("redis.DB"), viper.GetInt("redis.poolSize"),
		viper.GetInt("redis.minIdleConn"))
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(pong)
	}
}

func InitMysql() {
	err := mapper.InitMysql(viper.GetString("mysql.dns"))
	if err != nil {
		return
	}

	/*	user := models.UserBasic{}
		mapper.Open.Find(&user)
		fmt.Println(user)*/
}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	//fmt.Println(viper.Get("mysql"))
}
