package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "online-chat/docs"
	"online-chat/middleware"
	"online-chat/online_server/service"
)

func ApiRouter(router *gin.Engine) {
	router.Use(middleware.Translations())
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("index/", service.Index)
	userApi(router)
	wsApi(router)
}

func userApi(r *gin.Engine) {
	user := r.Group("user")
	user.POST("createUser", service.CreateUser)
	user.POST("login", service.Login)

	user.Use(middleware.AuthMiddleware())
	user.GET("getUserList", service.GetUserList)
	user.DELETE("deleteUser", service.DeleteUser)
	user.PUT("updateUser", service.UpdateUser)
	user.GET("getUser.name", service.GetUserByName)
	user.GET("getUser.id", service.GetUserByID)
	user.GET("getUser.phone", service.GetUserByPhone)
	user.GET("getUser.email", service.GetUserByEmail)
}

func wsApi(r *gin.Engine) {
	w := r.Group("ws")
	w.GET("sendMsg", service.SendMsgByWebSocket)
}
