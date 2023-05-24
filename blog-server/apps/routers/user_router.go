package routers

import (
	"blog-server/apps/api"
	"blog-server/middlerware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("My&28Hos)UngdskuUSNG(Njh"))

func (r *RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.Use(sessions.Sessions("sessionid", store))
	r.POST("email_login", userApi.EmailLoginView)
	r.GET("users", middlerware.JwtAuth(), userApi.UserListView)
	r.PUT("user_role", middlerware.JwtAdmin(), userApi.UserUpdateRoleView)
	r.PUT("user_pass", middlerware.JwtAuth(), userApi.UserUpdatePassView)
	r.GET("logout", middlerware.JwtAuth(), userApi.UserLogoutView)
	r.DELETE("users", middlerware.JwtAdmin(), userApi.UserDeleteList)
	r.POST("user_bind_email", middlerware.JwtAuth(), userApi.UserBindEmailView)
	r.POST("user_create", middlerware.JwtAdmin(), userApi.UserCreateView)
}
