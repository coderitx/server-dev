package routers

import "blog-server/api"

func (r *RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingApi
	r.GET("settings/:name", settingsApi.SettingsApiViews)
	r.PUT("settings/:name", settingsApi.SettingsApiUpdate)
}
