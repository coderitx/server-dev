package user_api

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/apps/service/common_svc"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sort"
)

// UserListView 用户菜单列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=responsex.ListResponse[models.UserModel]}
func (*UserApi) UserListView(c *gin.Context) {

	claimsVal, ok := c.Get("claims")
	if !ok {
		zap.S().Errorf("未登录")
		responsex.FailWithMessage("未登录", c)
		return
	}
	claims, ok := claimsVal.(*utils.CustomClaims)
	if !ok {
		zap.S().Errorf("解析token失败")
		responsex.FailWithMessage("token 错误", c)
		return
	}
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	list, count, _ := common_svc.ComList(&models.UserModel{}, common_svc.Options{
		PageInfo: common_svc.PageInfoValid(page),
	})
	var users []models.UserModel
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			user.UserName = ""
			user.Tel = utils.DesensitizationTel(user.Tel)
			user.Email = utils.DesensitizationEmail(user.Email)
		}
		users = append(users, *user)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
	responsex.OkWithList(users, count, c)
}
