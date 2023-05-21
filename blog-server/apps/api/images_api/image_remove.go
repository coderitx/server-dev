package images_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ImageDeleteList 删除图片
// @Tags 图片管理
// @Summary 批量删除图片
// @Description 批量删除图片
// @Param data body models.RemoveRequest true  "图片id列表"
// @Router /api/image [delete]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (i *ImagesApi) ImageDeleteList(c *gin.Context) {
	var req models.RemoveRequest
	c.ShouldBindJSON(&req)
	var imageList []models.BannerModel
	count := global.DB.Debug().Where("id in (?)", req.IDList).Find(&imageList).RowsAffected
	if count == 0 {
		responsex.FailWithMessage("图片不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	deleteData := []string{}
	for _, img := range imageList {
		deleteData = append(deleteData, img.Name)
	}
	responsex.Ok(deleteData, fmt.Sprintf("删除 %d 张图片", count), c)
	return
}
