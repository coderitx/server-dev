package middlerware

import (
	"blog-server/apps/models/ctype"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			responsex.FailWithMessage("未登录", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			zap.S().Errorf("解析token失败 [ERROR]: %v", err.Error())
			responsex.FailWithMessage("token 错误", c)
			c.Abort()
			return
		}
		res, err := global.RDB.Get(context.TODO(), token).Result()
		if res != "" && err == nil {
			zap.S().Errorf("已注销")
			responsex.FailWithMessage("已注销", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)

		c.Next()
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			responsex.FailWithMessage("未登录", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			zap.S().Errorf("解析token失败 [ERROR]: %v", err.Error())
			responsex.FailWithMessage("token 错误", c)
			c.Abort()
			return
		}
		res, err := global.RDB.Get(context.TODO(), token).Result()
		if res != "" && err == nil {
			zap.S().Errorf("已注销")
			responsex.FailWithMessage("已注销", c)
			c.Abort()
			return
		}
		if claims.Role != int(ctype.PermissionAdmin) {
			zap.S().Errorf("权限错误")
			responsex.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
