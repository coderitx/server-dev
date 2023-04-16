package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "online-chat/docs"
	"online-chat/online_server/service"
)

func ApiRouter(router *gin.Engine) {
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("index/", service.Index)
	userApi(router)
}

func userApi(r *gin.Engine) {
	user := r.Group("user")
	user.GET("get_user_list", service.GetUserList)
}
