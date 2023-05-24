package user_api

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserUpdateRoleRequest struct {
	UserID   uint       `json:"user_id" binding:"required" msg:"请输入用户ID"`
	NickName string     `json:"nick_name"`
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"请确认需要变更的权限是否正确"`
}

// UserUpdateRoleView 用户权限更新
// @Tags 用户管理
// @Summary 用户权限更新
// @Description 用户权限更新
// @Param data body UserUpdateRoleRequest    true  "用户权限参数"
// @Router /api/user_role [put]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*UserApi) UserUpdateRoleView(c *gin.Context) {
	var userRole UserUpdateRoleRequest
	err := c.ShouldBindJSON(&userRole)
	if err != nil {
		responsex.FailWithError(err, &userRole, c)
		return
	}
	fmt.Printf("%+v", userRole)
	var userModel models.UserModel
	err = global.DB.Take(&userModel, userRole.UserID).Error
	if err != nil {
		zap.S().Errorf("查询用户出现错误 [ERROR]: %v", err.Error())
		responsex.FailWithMessage("用户不存在", c)
		return
	}
	updates := make(map[string]any)
	updates["role"] = userRole.Role
	if userRole.NickName != "" {
		updates["nick_name"] = userRole.NickName
	}
	err = global.DB.Model(&userModel).Updates(updates).Error
	if err != nil {
		responsex.FailWithMessage("更新权限失败", c)
		return
	}
	responsex.OkWithData("更新权限成功", c)
	return
}
