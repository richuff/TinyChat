package service

import (
	"RcChat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Description 新增用户接口
// @Accept json
// @Produce json
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param repassword query string true "第二次输入的密码"
// @Success 200 {string} json{"code","message"}
// @Failure 400 {string} json{"code","message"}
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次输入的密码不一致",
		})
		return
	}
	user.Password = password
	user.LoginTime = time.Now()
	user.CreatedAt = time.Now()
	user.HeartBeatTime = time.Now()
	user.LoginOutTime = time.Now()
	models.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Description 删除用户接口
// @Accept json
// @Produce json
// @Param id query int true "用户id"
// @Success 200 {string} json{"code","message"}
// @Router /user/DeleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	user.ID = uint(id)
	models.DeleteUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
