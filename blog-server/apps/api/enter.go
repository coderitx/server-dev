package api

import (
	"blog-server/apps/api/images_api"
	"blog-server/apps/api/setting_api"
)

// ApiGroup
type ApiGroup struct {
	setting_api.SettingApi
	images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
