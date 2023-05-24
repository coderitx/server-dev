package routers

import (
	"blog-server/apps/api"
	"blog-server/middlerware"
)

func (r *RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.POST("email_login", userApi.EmailLoginView)
	r.GET("users", middlerware.JwtAuth(), userApi.UserListView)
	r.PUT("user_role", middlerware.JwtAdmin(), userApi.UserUpdateRoleView)
	r.PUT("user_pass", middlerware.JwtAuth(), userApi.UserUpdatePassView)
	r.GET("logout", middlerware.JwtAuth(), userApi.UserLogoutView)
}
