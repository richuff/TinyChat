package utils

import (
	"RcChat/mapper"
	"context"
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

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
func Publish(ctx context.Context, channel string, message string) error {
	var err error
	fmt.Println("Publish ", message)
	mapper.Red.Publish(ctx, channel, message)
	return err
}

// Subscribe 从redis订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	var err error
	sub := mapper.Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe ", msg.Payload)
	return msg.Payload, err
}
