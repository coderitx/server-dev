package setting_api

import (
	"blog-server/common/responsex"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SettingApi struct {
}

func (s *SettingApi) SettingsApiViews(c *gin.Context) {
	responsex.OkWithData(map[string]any{
		"code": http.StatusOK,
	}, c)
}
