package routers

import (
	"blog-server/apps/api"
	"blog-server/middlerware"
)

func (r *RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	r.Use(middlerware.JwtAuth())
	r.POST("tags", tagApi.TagCreateView)
	r.DELETE("tags", tagApi.TagDeleteView)
	r.PUT("tags/:id", tagApi.TagUpdateView)
	r.GET("tags", tagApi.TagListView)
}
