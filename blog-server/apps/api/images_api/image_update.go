package images_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

func (i *ImagesApi) ImageUpdateView(c *gin.Context) {
	var req ImageUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	var imgBanner models.BannerModel
	err = global.DB.Take(&imgBanner, req.ID).Error
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	err = global.DB.Model(&imgBanner).Update("name", req.Name).Error
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	responsex.OkWithData("图片更新成功", c)
	return
}
