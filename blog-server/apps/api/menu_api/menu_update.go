package menu_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuUpdateRequest struct {
}

// MenuUpdateView 图片更新
// @Tags 菜单管理
// @Summary 更新菜单
// @Description 更新菜单
// @Param data body MenuUpdateRequest    true  "广告的一些参数"
// @Router /api/menus/:id [put]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*MenuApi) MenuUpdateView(c *gin.Context) {
	var menuReq MenuRequest
	err := c.ShouldBindJSON(&menuReq)
	if err != nil {
		responsex.FailWithError(err, &menuReq, c)
		return
	}
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		responsex.FailWithMessage("菜单不存在", c)
		return
	}
	// 清空之前的banner信息
	err = global.DB.Model(&menuModel).Association("Banners").Clear()
	if err != nil {
		zap.S().Errorf("更新菜单图片失败 [ERROR]: %v", err.Error())
		responsex.FailWithMessage("更新菜单图片失败", c)
		return
	}
	// 存在菜单图片信息则重新添加
	if len(menuReq.ImageSortList) != 0 {
		// 操作关联表
		var bannerMenu []models.MenuBannerModel
		for _, s := range menuReq.ImageSortList {
			bannerMenu = append(bannerMenu, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: s.ImageID,
				Sort:     s.Sort,
			})
		}
		err = global.DB.Create(&bannerMenu).Error
		if err != nil {
			zap.S().Errorf("创建菜单图片失败 [ERROR]: %v", err.Error())
			responsex.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	updates := structs.Map(&menuReq)
	err = global.DB.Model(&menuModel).Updates(updates).Error
	if err != nil {
		zap.S().Errorf("更新菜单失败失败 [ERROR]: %v", err.Error())
		responsex.FailWithMessage("更新菜单失败", c)
		return
	}
	responsex.OkWithData("更新菜单成功", c)
	return
}
