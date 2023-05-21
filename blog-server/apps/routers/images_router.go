package routers

import (
	"blog-server/apps/api"
)

func (r *RouterGroup) ImagesRouter() {
	imageApi := api.ApiGroupApp.ImagesApi
	r.POST("upload", imageApi.ImageUploadView)
	r.GET("image", imageApi.ImageListView)
	r.GET("image_name", imageApi.ImageNameListView)
	r.DELETE("image", imageApi.ImageDeleteList)
	r.PUT("image", imageApi.ImageUpdateView)
}
