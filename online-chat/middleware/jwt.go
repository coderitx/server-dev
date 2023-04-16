package middleware

import (
	"github.com/gin-gonic/gin"
	"online-chat/common/response"
	"online-chat/utils"
	"strconv"
	"time"
)

var ExpiresAt int64 = 300

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		auth := c.Request.Header.Get("token")
		if auth == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		claims, err := utils.AnalyzeToken(auth)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "登陆验证失败或登陆信息已过期", c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < ExpiresAt {
			claims.ExpiresAt = time.Now().Unix() + ExpiresAt
			newToken, _ := utils.GenerateToken(claims.Id, claims.Identity, claims.Name, claims.Phone, claims.Email, 300)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(ExpiresAt, 10))

		}
		c.Set("claims", claims)
		c.Next()
	}
}
