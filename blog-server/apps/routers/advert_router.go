package routers

import "blog-server/apps/api"

func (r *RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	r.POST("/advert", advertApi.AdvertCreateView)
	r.GET("/advert", advertApi.AdvertListView)
	r.PUT("/advert/:id", advertApi.AdvertUpdateView)
	r.DELETE("/advert", advertApi.AdvertDeleteView)
}
