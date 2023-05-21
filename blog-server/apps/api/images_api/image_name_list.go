package images_api

import (
	"blog-server/apps/models"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageNameListResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameListView 查询图片名称信息列表
// @Tags 图片管理
// @Summary 图片名称信息列表
// @Description 图片名称信息列表
// @Router /api/image_name [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=[]ImageNameListResponse}
func (*ImagesApi) ImageNameListView(c *gin.Context) {
	var images []ImageNameListResponse
	err := global.DB.Model(&models.BannerModel{}).Select("id", "name", "path").Scan(&images).Error
	if err != nil {
		zap.S().Error("query image list error: %v", err.Error())
		responsex.FailWithCode(errorx.SettingsError, c)
		return
	}
	responsex.OkWithData(images, c)
	return
}
