package tag_api

import (
	"blog-server/apps/models"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest    true  "标签的一些参数"
// @Router /api/tags/:id [put]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var tagReq TagRequest
	err := c.ShouldBindJSON(&tagReq)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		zap.S().Error("id = %v 的标签不存在", id)
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	count := global.DB.Take(&models.TagModel{}, "title = ?", tagReq.Title).RowsAffected
	if count != 0 {
		responsex.Fail(tagReq, "更新的title重复,请重新输入标题", c)
		return
	}
	err = global.DB.Model(&tag).Updates(structs.Map(&tagReq)).Debug().Error
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	responsex.OkWithData("更新成功", c)
	return
}
