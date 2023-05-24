package user_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserDeleteList 删除用户
// @Tags 用户管理
// @Summary 批量删除用户
// @Description 批量删除用户
// @Param data body models.RemoveRequest true  "用户id列表"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (i *UserApi) UserDeleteList(c *gin.Context) {
	var req models.RemoveRequest
	c.ShouldBindJSON(&req)
	var imageList []models.BannerModel
	count := global.DB.Debug().Where("id in (?)", req.IDList).Find(&imageList).RowsAffected
	if count == 0 {
		responsex.FailWithMessage("图片不存在", c)
		return
	}
	// TODO: 用户关联相关表操作删除
	tx := global.DB.Begin()
	tx.Delete(&imageList)
	deleteData := []string{}
	for _, img := range imageList {
		deleteData = append(deleteData, img.Name)
	}
	responsex.Ok(deleteData, fmt.Sprintf("删除 %d 个用户", count), c)
	return
}
