package user_api

import (
	"blog-server/apps/models"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
)

type UpdatePassRequest struct {
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

// UserUpdatePassView 用户修改密码
// @Tags 用户管理
// @Summary 用户密码更新
// @Description 用户密码更新
// @Param data body UpdatePassRequest    true  "用户密码更新参数"
// @Router /api/user_pass [put]
// @Produce json
// @Success 200 {object} responsex.Response{data=any}
func (*UserApi) UserUpdatePassView(c *gin.Context) {
	var updatePass UpdatePassRequest
	err := c.ShouldBindJSON(&updatePass)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, ok := _claims.(*utils.CustomClaims)
	if !ok {
		responsex.FailWithMessage("登陆信息错误", c)
		return
	}
	userId := claims.UserID
	oldPassMd5 := utils.MD5([]byte(updatePass.OldPass))
	newPassMd5 := utils.MD5([]byte(updatePass.NewPass))
	err = global.DB.Take(&models.UserModel{}, userId).Error
	if err != nil {
		responsex.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Where("id = ? AND password = ?", userId, oldPassMd5).Take(&models.UserModel{}).Error
	if err != nil {
		responsex.FailWithMessage("旧密码输入错误", c)
		return
	}
	err = global.DB.Model(&models.UserModel{}).Where("id = ?", userId).Update("password", newPassMd5).Error
	if err != nil {
		responsex.FailWithMessage("更新密码失败", c)
		return
	}
	responsex.OkWithData("更新密码成功", c)
	return
}
