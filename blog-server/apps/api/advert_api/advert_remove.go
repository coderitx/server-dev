package advert_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertDeleteView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest true  "广告id列表"
// @Router /api/advert [delete]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*AdvertApi) AdvertDeleteView(c *gin.Context) {
	var req models.RemoveRequest
	c.ShouldBindJSON(&req)
	var advertList []models.AdvertModel
	count := global.DB.Debug().Where("id in (?)", req.IDList).Find(&advertList).RowsAffected
	if count == 0 {
		responsex.FailWithMessage("图片不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	deleteData := []string{}
	for _, adv := range advertList {
		deleteData = append(deleteData, adv.Title)
	}
	responsex.Ok(deleteData, fmt.Sprintf("删除 %d 条广告", count), c)
	return
}
