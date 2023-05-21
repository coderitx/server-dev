package advert_api

import (
	"blog-server/apps/models"
	"blog-server/apps/service/common_svc"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"github.com/gin-gonic/gin"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Produce json
// @Param data query models.PageInfo false "查询参数"
// @Success 200 {object} responsex.Response{data=responsex.ListResponse[models.AdvertModel]}
// @Failure 400 {object} responsex.Response{}
// @Failure 500 {object} responsex.Response{}
// @Router /api/advertList [get]
func (*AdvertApi) AdvertListView(c *gin.Context) {
	var pageInfo models.PageInfo
	err := c.ShouldBindUri(&pageInfo)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	// 根据请求头Referer判断请求来源，前台还是后台管理页面
	advertModel := models.AdvertModel{}
	referer := c.GetHeader("Referer")
	if !strings.Contains(referer, "admin") {
		advertModel.IsShow = true
	}
	// gorm 特效，如果isShow==false，但是isShow=true的也会查询到
	list, count, _ := common_svc.ComList(advertModel, common_svc.Options{
		PageInfo: common_svc.PageInfoValid(pageInfo),
	})
	responsex.OkWithList(list, count, c)
	return
}
