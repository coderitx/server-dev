package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"online-chat/common/errorx"
	"online-chat/common/response"
	"online-chat/online_server/models/wsx"
)

var upgrade = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMsgByWebSocket 通过websocket发送消息
// @Router /ws/sendMsg [get]
func SendMsgByWebSocket(ctx *gin.Context) {
	m := wsx.WSMessage{}
	code := m.SendMsg(ctx, upgrade)
	response.Failed(ctx, http.StatusBadRequest, code, errorx.ErrMsg(code), nil)
	return
}
