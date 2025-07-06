package router

import (
	"RcChat/docs"
	"RcChat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Static("/favicon.ico", "/common/favicon.ico")
	r.Static("/common/file", "/common/file")

	SwaggerRouter(r)
	IndexRouter(r)
	UserRouter(r)
	ContactRouter(r)

	r.GET("/attach/upload", service.UploadFile)
	return r
}

func IndexRouter(r *gin.Engine) {
	//index
	r.GET("/index", service.GetIndex)
}

func SwaggerRouter(r *gin.Engine) {
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func UserRouter(r *gin.Engine) {
	userApi := r.Group("/user")
	userApi.GET("/GetUserList", service.GetUserList)
	userApi.GET("/CreateUser", service.CreateUser)
	userApi.GET("/DeleteUser", service.DeleteUser)
	userApi.POST("/UpdateUser", service.UpdateUser)
	userApi.POST("/UserLogin", service.UserLogin)
	userApi.GET("/SendMsg", service.SendMessage)
	userApi.GET("/SendUserMsg", service.SendMsg)
	userApi.GET("/searchFriends", service.SearchFriend)
}

func ContactRouter(r *gin.Engine) {
	contactApi := r.Group("/contact")
	contactApi.POST("/addFriend", service.AddFriend)
	contactApi.GET("/addCommunity", service.CreateCommunity)
	contactApi.GET("/loadCommunity", service.LoadCommunity)
	contactApi.GET("/joinCommunity", service.JoinCommunity)
}
