package service

import (
	"github.com/gin-gonic/gin"
)

// Index
// @Router /index [get]
func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello online-chat",
	})
}
