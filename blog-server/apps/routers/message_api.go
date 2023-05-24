package routers

import "blog-server/apps/api"

func (r *RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	r.POST("messages", messageApi.MessageCreateView)
	r.GET("message_all", messageApi.MessageListAllView)
	r.GET("messages", messageApi.MessageListView)
	r.GET("messages_record", messageApi.MessageRecordView)
}
