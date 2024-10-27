package mapper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Open *gorm.DB
)

func InitMysql(config string) (err error) {
	//连接数据库
	Open, err = gorm.Open("mysql", config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return Open.DB().Ping()
}
