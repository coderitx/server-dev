package tag_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"标题不合法" structs:"title"` // 显示标题
}

// TagCreateView 增加标签
// @Tags 标签管理
// @Summary 增加标签
// @Description 增加标签
// @Produce json
// @Param data body TagRequest true "表示多个参数"
// @Success 200 {object} responsex.Response{data=string}
// @Router /api/tags [post]
func (*TagApi) TagCreateView(c *gin.Context) {
	var req TagRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	// 判断重复广告
	var advert models.TagModel
	err = global.DB.Take(&advert, "title = ?", req.Title).Error
	if err == nil {
		zap.S().Error("标签已存在")
		responsex.FailWithMessage("标签已存在", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: req.Title,
	}).Debug().Error
	if err != nil {
		zap.S().Error("添加标签失败[ERROR]: %v", err.Error())
		responsex.FailWithMessage("添加标签失败", c)
		return
	}
	responsex.OkWithData("添加标签成功", c)
	return
}
