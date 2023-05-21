package menu_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Banner struct {
	ID   uint   `json:"image_id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 查询菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=MenuResponse[]}
func (*MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	err := global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList).Debug().Error
	if err != nil {
		zap.S().Errorf("菜单不存在[ERROR]: %v", err.Error())
		responsex.FailWithMessage("菜单不存在", c)
		return
	}

	var menuBanners []models.MenuBannerModel
	// 预加载关联表信息
	err = global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners).Where("menu_id in (?)", menuIDList).Debug().Error
	if err != nil {
		zap.S().Errorf("查询菜单关联图片失败[ERROR]: %v", err.Error())
		responsex.FailWithMessage("查询菜单关联图片失败", c)
		return
	}
	var menuRes []MenuResponse
	// 遍历所有菜单,封装每一个菜单的图片信息
	for _, menuMod := range menuList {
		// 所有菜单关联的图片
		banners := []Banner{}
		for _, b := range menuBanners {
			if menuMod.ID != b.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   b.BannerModel.ID,
				Path: b.BannerModel.Path,
			})
		}
		menuRes = append(menuRes, MenuResponse{
			MenuModel: menuMod,
			Banners:   banners,
		})
	}

	responsex.OkWithData(menuRes, c)
	return
}
