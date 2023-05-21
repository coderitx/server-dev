package menu_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MenuDetailsView 查询菜单列表
// @Tags 菜单管理
// @Summary 菜单详情
// @Description 菜单详情
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=MenuResponse}
func (*MenuApi) MenuDetailsView(c *gin.Context) {
	id := c.Param("id")
	// 查询出菜单表的信息
	menuModel := models.MenuModel{}
	err := global.DB.Find(&menuModel, id).Error
	if err != nil {
		zap.S().Errorf("菜单不存在 [ERROR]: %v", err.Error())
		responsex.FailWithMessage("菜单不存在", c)
		return
	}

	// 查询出中间表图片的信息
	var menuBannerModel []models.MenuBannerModel
	err = global.DB.Preload("BannerModel").Find(&menuBannerModel, "menu_id = ?", id).Error
	if err != nil {
		zap.S().Errorf("查询菜单关联图片失败[ERROR]: %v", err.Error())
		responsex.FailWithMessage("查询菜单关联图片失败", c)
		return
	}

	// 所有菜单关联的图片
	banners := []Banner{}
	for _, b := range menuBannerModel {
		// 如果当前菜单不存在图片信息，则直接跳过
		if menuModel.ID != b.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   b.BannerModel.ID,
			Path: b.BannerModel.Path,
		})
	}
	menuRes := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	if menuRes.Title == "" {
		responsex.OkWithData(nil, c)
		return
	}
	responsex.OkWithData(menuRes, c)
	return
}
