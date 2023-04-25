package api

import "blog-server/api/setting_api"

// ApiGroup
type ApiGroup struct {
	setting_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
