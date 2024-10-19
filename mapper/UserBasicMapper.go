package mapper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Open *gorm.DB
)

func InitMysql() (err error) {
	//连接数据库
	Open, err = gorm.Open("mysql", "root:password@(127.0.0.1:3306)/mp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return Open.DB().Ping()
}
