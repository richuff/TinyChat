package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ErrHandler(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(200, gin.H{
			"message": "输入格式有问题",
		})
		log.Println(err)
		return true
	}
	return false
}
