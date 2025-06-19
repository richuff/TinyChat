package service

import (
	"RcChat/constant"
	"RcChat/mapper"
	"RcChat/models"
	"RcChat/utils"
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var redisSub *redis.PubSub

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
		result := constant.NewResult().SetCode(1).SetMessage("该用户名已注册")
		result.Success(c)
		return
	}

	if password != repassword {
		result := constant.NewResult().SetCode(1).SetMessage("两次输入的密码不一致")
		result.Success(c)
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
	result := constant.NewResult().SetCode(1).SetMessage("创建成功")
	result.Success(c)
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
	log.Println(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		log.Println(err)
		result := constant.NewResult().SetCode(0).SetMessage("邮箱或手机号格式错误")
		result.Error(c)
		return
	}
	models.UpdateUser(user)
	result := constant.NewResult().SetCode(1).SetMessage("注册成功")
	result.SuccessByData(c, user)
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

// UpGrade 防止跨域站点的伪请求
var UpGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMessage 发送消息
//func SendMessage(c *gin.Context) {
//	ws, err := UpGrade.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	MsgHandler(ws, c)
//}
//func MsgHandler(ws *websocket.Conn, c *gin.Context) {
//	for {
//		msg, err := utils.Subscribe(c, utils.PublishKey)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		tm := time.Now().Format("2006-01-02 15:04:05")
//		m := fmt.Sprintf("[ws][%s]:[%s]", tm, msg)
//		fmt.Println(m)
//		err = ws.WriteMessage(websocket.TextMessage, []byte(m))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}

// SendMessage 处理 WebSocket 升级并启动消息处理、
func SendMessage(c *gin.Context) {
	redisSub = mapper.Red.Subscribe(context.Background(), utils.PublishKey)
	ws, err := UpGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Println(err)
		}
	}(ws)

	go func() {
		for {
			msg, err := redisSub.ReceiveMessage(context.Background())
			if err != nil {
				log.Println("Receive message from Redis error:", err)
				return
			}
			tm := time.Now().Format("2006-01-02 15:04:05")
			m := fmt.Sprintf("[ws][%s]:[%s]", tm, msg.Payload)
			log.Println(m)
			err = ws.WriteMessage(websocket.TextMessage, []byte(m))
			if err != nil {
				log.Println("Write message to WebSocket error:", err)
				return
			}
		}
	}()

	// 处理从客户端接收的消息
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read message from WebSocket error:", err)
			return
		}
		message := string(p)
		log.Println("Received message from client:", message)

		// 将消息发布到 Redis 频道，以便其他客户端可以接收
		err = mapper.Red.Publish(context.Background(), utils.PublishKey, message).Err()
		if err != nil {
			log.Println("Publish message to Redis error:", err)
			return
		}
	}
}

func SendMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriend(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("userId"), 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	users := models.SearchFriend(userId)
	utils.RespOkList(c.Writer, users, len(users))
}
