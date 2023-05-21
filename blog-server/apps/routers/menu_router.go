package routers

import "blog-server/apps/api"

func (r *RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	r.POST("menus", menuApi.MenuCreateView)
	r.GET("menus", menuApi.MenuListView)
	r.GET("menus_name", menuApi.MenuNameListView)
	r.PUT("menus/:id", menuApi.MenuUpdateView)
	r.DELETE("menus", menuApi.MenuDeleteList)
	r.GET("menus/:id", menuApi.MenuDetailsView)
}
