package tag_api

import (
	"blog-server/apps/models"
	"blog-server/apps/service/common_svc"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"github.com/gin-gonic/gin"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Produce json
// @Param data query models.PageInfo false "查询参数"
// @Success 200 {object} responsex.Response{data=responsex.ListResponse[models.TagModel]}
// @Router /api/tags [get]
func (*TagApi) TagListView(c *gin.Context) {
	var pageInfo models.PageInfo
	err := c.ShouldBindUri(&pageInfo)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	tagModel := models.TagModel{}
	list, count, _ := common_svc.ComList(tagModel, common_svc.Options{
		PageInfo: common_svc.PageInfoValid(pageInfo),
	})
	responsex.OkWithList(list, count, c)
	return
}
