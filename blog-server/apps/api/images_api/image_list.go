package images_api

import (
	"blog-server/apps/models"
	"blog-server/apps/service/common_svc"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ImageListView 查询图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/image [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=responsex.ListResponse[models.BannerModel]}
func (i *ImagesApi) ImageListView(c *gin.Context) {
	var p models.PageInfo
	err := c.ShouldBindQuery(&p)
	if err != nil {
		zap.S().Errorf("shoud bind params error: %v", err.Error())
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	list, count, err := common_svc.ComList(models.BannerModel{}, common_svc.Options{
		PageInfo: common_svc.PageInfoValid(p),
	})
	if err != nil {
		zap.S().Error("query image list error: %v", err.Error())
		responsex.FailWithCode(errorx.SettingsError, c)
		return
	}
	responsex.OkWithList(list, count, c)
	return
}
