package routers

import "blog-server/apps/api"

func (r *RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	r.POST("/advertCreate", advertApi.AdvertCreateView)
	r.GET("/advertList", advertApi.AdvertListView)
	r.PUT("/advertUpdate/:id", advertApi.AdvertUpdateView)
	r.DELETE("/advertDelete", advertApi.AdvertDeleteView)
}
