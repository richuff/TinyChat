package mapper

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Open  *gorm.DB
	SqlDB *sql.DB
	Red   *redis.Client
)

func InitRedis(addr string, password string, db int, poolSize int, minIdleConns int) (string, error) {
	Red = redis.NewClient(&redis.Options{
		Addr:         addr,         // Redis 服务器地址和端口
		Password:     password,     // Redis 访问密码，如果没有可以为空字符串
		DB:           db,           // 使用的 Redis 数据库编号，默认为 0
		PoolSize:     poolSize,     //指定连接池中的最大连接数
		MinIdleConns: minIdleConns, //指定连接池中的最小空闲连接数
	})
	pong, err := Red.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
	}
	return pong, err
}

func InitMysql(config string) (err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)
	//连接数据库
	Open, err = gorm.Open(mysql.Open(config), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	SqlDB, err = Open.DB()

	return err
}
