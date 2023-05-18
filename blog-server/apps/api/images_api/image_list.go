package images_api

import (
	"blog-server/apps/models"
	"blog-server/apps/service/svc_common"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ImageListView 查询图片列表
func (i *ImagesApi) ImageListView(c *gin.Context) {
	var p models.PageInfo
	err := c.ShouldBindQuery(&p)
	if err != nil {
		zap.S().Errorf("shoud bind params error: %v", err.Error())
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	list, count, err := svc_common.ComList(models.BannerModel{}, svc_common.Options{
		PageInfo: p,
	})
	if err != nil {
		zap.S().Error("query image list error: %v", err.Error())
		responsex.FailWithCode(errorx.SettingsError, err.Error(), c)
		return
	}
	responsex.OkWithList(list, count, c)
	return
}
