package service

import (
	"RcChat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList
// @Tags 首页
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
