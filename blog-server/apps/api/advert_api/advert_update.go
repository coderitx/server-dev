package advert_api

import (
	"blog-server/apps/models"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Router /api/advert/:id [put]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var advertReq AdvertRequest
	err := c.ShouldBindJSON(&advertReq)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	var adv models.AdvertModel
	err = global.DB.Take(&adv, id).Error
	if err != nil {
		zap.S().Error("id = %v 的广告不存在", id)
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	count := global.DB.Take(&models.AdvertModel{}, "title = ?", advertReq.Title).RowsAffected
	if count != 0 {
		responsex.Fail(advertReq, "更新的title重复,请重新输入标题", c)
		return
	}
	err = global.DB.Model(&adv).Updates(structs.Map(&advertReq)).Debug().Error
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	responsex.OkWithData("更新成功", c)
	return
}
