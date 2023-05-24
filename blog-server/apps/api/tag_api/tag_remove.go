package tag_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

// TagDeleteView 批量删除标签
// @Tags 标签管理
// @Summary 批量删除标签
// @Description 批量删除标签
// @Param data body models.RemoveRequest true  "标签id列表"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*TagApi) TagDeleteView(c *gin.Context) {
	var req models.RemoveRequest
	c.ShouldBindJSON(&req)
	var advertList []models.TagModel
	count := global.DB.Debug().Where("id in (?)", req.IDList).Find(&advertList).RowsAffected
	if count == 0 {
		responsex.FailWithMessage("标签不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	deleteData := []string{}
	for _, adv := range advertList {
		deleteData = append(deleteData, adv.Title)
	}
	responsex.Ok(deleteData, fmt.Sprintf("删除 %d 条标签", count), c)
	return
}
