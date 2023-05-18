package routers

import (
	"blog-server/apps/api"
)

func (r *RouterGroup) ImagesRouter() {
	imageApi := api.ApiGroupApp.ImagesApi
	r.POST("upload", imageApi.ImageUploadView)
	r.GET("imageList", imageApi.ImageListView)
	r.DELETE("imageDelete", imageApi.ImageDeleteList)
	r.PUT("imageUpdate", imageApi.ImageUpdateView)
}
