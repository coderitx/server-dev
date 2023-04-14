package router

import "github.com/gin-gonic/gin"

func ApiRouter(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hell chat",
		})
	})
}
