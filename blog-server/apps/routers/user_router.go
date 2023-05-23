package routers

import "blog-server/apps/api"

func (r *RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.POST("email_login", userApi.EmailLoginView)
	r.GET("users", userApi.UserListView)
}
