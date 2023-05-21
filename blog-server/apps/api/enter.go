package api

import (
	"blog-server/apps/api/advert_api"
	"blog-server/apps/api/images_api"
	"blog-server/apps/api/menu_api"
	"blog-server/apps/api/setting_api"
)

// ApiGroup
type ApiGroup struct {
	setting_api.SettingApi
	images_api.ImagesApi
	advert_api.AdvertApi
	menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
