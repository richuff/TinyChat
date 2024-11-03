package service

import (
	"RcChat/models"
	"RcChat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/GetUserList [get]
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

	slat := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "该用户名已注册",
		})
		return
	}

	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次输入的密码不一致",
		})
		return
	}
	/*user.Password = password*/
	user.Salt = slat
	user.Password = utils.MakePassword(password, slat)
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

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @Description 更新用户接口
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Param email formData string false "email"
// @Param phone formData string false "phone"
// @Success 200 {string} json{"code","message"}
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	fmt.Println(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "格式错误",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// UserLogin
// @Summary 用户登录
// @Tags 用户模块
// @Description 用户登录接口
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/UserLogin [post]
func UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	data := models.FindUserByName(name)
	if data.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "该用户还未注册",
		})
		return
	}
	slat := data.Salt
	if data.Password != utils.ValidPassword(password, slat) {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码或用户名错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})
}
