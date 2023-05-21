package menu_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// MenuNameListView 查询菜单名称信息列表
// @Tags 菜单管理
// @Summary 菜单名称信息列表
// @Description 菜单名称信息列表
// @Router /api/menus_name [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=[]MenuNameResponse}
func (*MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameResponse []MenuNameResponse
	err := global.DB.Model(&models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameResponse).Debug().Error
	if err != nil {
		zap.S().Errorf("获取菜单名称信息失败 [ERROR: %v]", err.Error())
		responsex.FailWithMessage("获取菜单名称信息失败", c)
		return
	}
	responsex.OkWithData(menuNameResponse, c)
	return
}
