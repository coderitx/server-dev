package user_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView 邮箱登陆
// @Tags 用户管理
// @Summary 邮箱登陆
// @Description 邮箱登陆
// @Produce json
// @Param data body LoginRequest true "表示多个参数"
// @Success 200 {object} responsex.Response{data=string}
// @Router /api/email_login [post]
func (*UserApi) EmailLoginView(c *gin.Context) {
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", req.UserName, req.UserName).Error
	if err != nil {
		zap.S().Errorf("用户名不存在 [ERROR]: %v", err.Error())
		responsex.FailWithMessage("用户名不存在", c)
		return
	}
	pass := utils.MD5([]byte(req.Password))
	if userModel.Password != pass {
		responsex.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 登陆成功,生成token
	token, _ := utils.GenerateToken(utils.JWTPayload{
		NickName: userModel.NickName,
		UserID:   userModel.ID,
		Role:     int(userModel.Role),
	})
	responsex.OkWithData(map[string]string{
		"token": token,
	}, c)
	return
}
