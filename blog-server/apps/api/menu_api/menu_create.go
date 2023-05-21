package menu_api

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`            // 菜单图片的切换时间 单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                // 菜单图片的切换时间 单位秒
	Sort          int         `gorm:"size:10" json:"sort" msg:"请输入菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
}

// MenuCreateView 添加菜单
// @Tags 菜单管理
// @Summary 创建菜单
// @Description 创建菜单
// @Produce json
// @Param data body MenuRequest true "表示多个参数"
// @Success 200 {object} responsex.Response{data=string}
// @Router /api/menus [post]
func (*MenuApi) MenuCreateView(c *gin.Context) {
	var req MenuRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	//判断是否已经存在
	var menuModel models.MenuModel
	err = global.DB.Where("title = ? or path = ?", req.Title, req.Path).Find(&menuModel).Error
	if err != nil {
		zap.S().Errorf("检查是否存在发生错误 [ERROR]: %v", err.Error())
		responsex.FailWithMessage(fmt.Sprintf("title = %v 的菜单已经存在", req.Title), c)
		return
	}
	newMenuModel := models.MenuModel{
		Title:        req.Title,
		Path:         req.Path,
		Slogan:       req.Slogan,
		Abstract:     req.Abstract,
		AbstractTime: req.AbstractTime,
		BannerTime:   req.BannerTime,
		Sort:         req.Sort,
	}

	// 菜单基本信息入库
	err = global.DB.Create(&newMenuModel).Error
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	if len(req.ImageSortList) == 0 {
		responsex.OkWithData("添加菜单成功", c)
		return
	}
	// 判断图片列表id的图片是否存在
	var reqImageId []uint
	for _, img := range req.ImageSortList {
		reqImageId = append(reqImageId, img.ImageID)
	}
	var imageID []uint
	global.DB.Model(&models.BannerModel{}).Where("id in (?)", reqImageId).Pluck("id", &imageID)

	var menuBannerModel []models.MenuBannerModel
	for _, mb := range req.ImageSortList {
		if !IN(mb.ImageID, imageID) {
			continue
		}
		menuBannerModel = append(menuBannerModel, models.MenuBannerModel{
			MenuID:   newMenuModel.ID,
			BannerID: mb.ImageID,
			Sort:     mb.Sort,
		})
	}
	// 菜单关联图片入库
	err = global.DB.Create(&menuBannerModel).Error
	if err != nil {
		zap.S().Error("菜单图片关联失败")
		responsex.OkWithData("添加菜单图片失败", c)
		return
	}
	responsex.OkWithData("添加菜单成功", c)
	return
}

func IN(id uint, idList []uint) bool {
	for _, newId := range idList {
		if id == newId {
			return true
		}
	}
	return false
}
