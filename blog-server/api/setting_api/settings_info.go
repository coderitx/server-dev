package setting_api

import (
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
)

func (s *SettingApi) SettingsApiViews(c *gin.Context) {
	var su SettingUri
	err := c.ShouldBindUri(&su)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError, err.Error(), c)
		return
	}
	switch su.Name {
	case "site":
		responsex.OkWithData(global.GlobalC.SiteInfo, c)
	case "email":
		responsex.OkWithData(global.GlobalC.Email, c)
	case "qq":
		responsex.OkWithData(global.GlobalC.QQ, c)
	case "tencent":
		responsex.OkWithData(global.GlobalC.Tencent, c)
	default:
		responsex.FailWithMessage("没有对应的配置信息", c)
	}
}
