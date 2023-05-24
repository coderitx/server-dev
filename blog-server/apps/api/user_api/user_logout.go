package user_api

import (
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// UserLogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Router /api/logout [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=any}
func (*UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims, ok := _claims.(*utils.CustomClaims)
	if !ok {
		responsex.FailWithMessage("登陆信息有误", c)
		return
	}
	token := c.Request.Header.Get("token")
	exp := claims.ExpiresAt.Sub(time.Now())
	err := global.RDB.Set(context.TODO(), token, "logout", exp).Err()
	if err != nil {
		responsex.FailWithMessage("注销失败", c)
		return
	}
	responsex.OkWithData("注销成功", c)
	return

}
