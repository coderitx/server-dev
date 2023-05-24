package user_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/plugins/email"
	"blog-server/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserBindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱不合法"`
	Password string  `json:"password" msg:"请输入密码"`
	Code     *string `json:"code"`
}

func (*UserApi) UserBindEmailView(c *gin.Context) {
	claimsVal, ok := c.Get("claims")
	if !ok {
		zap.S().Errorf("未登录")
		responsex.FailWithMessage("未登录", c)
		return
	}
	claims, ok := claimsVal.(*utils.CustomClaims)
	// 1. 用户输入邮箱
	var req UserBindEmailRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	session := sessions.Default(c)
	if req.Code == nil {
		// 2. 发送验证码
		code := utils.RandCode(6)
		// 2.1 写入session
		session.Set("valid_code", code)
		err := email.NewCode().Send(req.Email, code)
		if err != nil {
			responsex.FailWithMessage("验证码发送失败", c)
			return
		}
		session.Save()
		responsex.OkWithData("验证码已发送成功，请查收", c)
		return
	}
	// 3. 输入邮箱验证码，密码，验证之后绑定
	code := session.Get("valid_code")
	if code != *req.Code {
		responsex.FailWithMessage("验证码不正确", c)
		return
	}
	// 4. 修改用户邮箱
	var userModel models.UserModel
	err = global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		responsex.FailWithMessage("用户不存在", c)
		return
	}
	if len(req.Password) < 4 {
		responsex.FailWithMessage("密码强度太低", c)
		return
	}
	passMd5 := utils.MD5([]byte(req.Password))
	err = global.DB.Model(&userModel).Updates(map[string]any{
		"email":    req.Email,
		"password": passMd5,
	}).Error
	if err != nil {
		zap.S().Errorf("绑定邮箱错误 [ERROR]: %v", err.Error())
		responsex.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 5. 返回成功
	responsex.OkWithData("绑定邮箱成功", c)
	return
}
