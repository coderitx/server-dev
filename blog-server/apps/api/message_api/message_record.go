package message_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id"`
}

// MessageRecordView 消息聊天列表
// @Tags 消息管理
// @Summary 消息聊天列表
// @Description 消息聊天列表
// @Router /api/messages_record [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*MessageApi) MessageRecordView(c *gin.Context) {
	var req MessageRecordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("created_at asc").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range _messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		if model.RevUserID == req.UserID || model.SendUserID == req.UserID {
			messageList = append(messageList, model)
		}
	}

	// 点开消息，里面的每一条消息，都从未读变成已读

	responsex.OkWithData(messageList, c)
	return
}
