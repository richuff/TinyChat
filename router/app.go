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
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", service.GetIndex)
	r.GET("/user/GetUserList", service.GetUserList)
	r.GET("/user/CreateUser", service.CreateUser)
	r.GET("/user/DeleteUser", service.DeleteUser)
	r.POST("/user/UpdateUser", service.UpdateUser)
	r.POST("/user/UserLogin", service.UserLogin)

	r.GET("/user/SendMsg", service.SendMessage)
	r.GET("/user/SendUserMsg", service.SendMsg)
	return r
}
