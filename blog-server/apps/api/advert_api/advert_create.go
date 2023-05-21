package advert_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"标题不合法" structs:"title"`      // 显示标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接不合法" structs:"href"'` // 跳转链接
	Images string `json:"image" binding:"required" msg:"图片地址不合法" structs:"images"`   // 图片
	IsShow bool   `json:"is_show" structs:"is_show"`                                 // 是否展示
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Produce json
// @Param data body AdvertRequest true "表示多个参数"
// @Success 200 {object} responsex.Response{data=string}
// @Router /api/advert [post]
func (*AdvertApi) AdvertCreateView(c *gin.Context) {
	var req AdvertRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	// 判断重复广告
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", req.Title).Error
	if err == nil {
		zap.S().Error("广告已存在")
		responsex.FailWithMessage("广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  req.Title,
		Href:   req.Href,
		Images: req.Images,
		IsShow: req.IsShow,
	}).Debug().Error
	if err != nil {
		zap.S().Error("添加广告失败[ERROR]: %v", err.Error())
		responsex.FailWithMessage("添加广告失败", c)
		return
	}
	responsex.OkWithData("添加广告成功", c)
	return
}
