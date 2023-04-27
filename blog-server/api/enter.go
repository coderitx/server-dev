package api

import (
	"blog-server/api/images_api"
	"blog-server/api/setting_api"
)

// ApiGroup
type ApiGroup struct {
	setting_api.SettingApi
	images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
