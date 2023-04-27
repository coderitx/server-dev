package routers

import "blog-server/api"

func (r *RouterGroup) ImagesRouter() {
	imageApi := api.ApiGroupApp.ImagesApi
	r.POST("upload", imageApi.ImagesUploadView)
}
