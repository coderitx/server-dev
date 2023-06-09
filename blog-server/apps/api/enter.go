package api

import (
	"blog-server/apps/api/advert_api"
	"blog-server/apps/api/images_api"
	"blog-server/apps/api/menu_api"
	"blog-server/apps/api/message_api"
	"blog-server/apps/api/setting_api"
	"blog-server/apps/api/tag_api"
	"blog-server/apps/api/user_api"
)

// ApiGroup
type ApiGroup struct {
	setting_api.SettingApi
	images_api.ImagesApi
	advert_api.AdvertApi
	menu_api.MenuApi
	user_api.UserApi
	tag_api.TagApi
	message_api.MessageApi
}

var ApiGroupApp = new(ApiGroup)
