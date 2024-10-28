package mapper

import (
	"database/sql"
	"fmt"
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
)

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
