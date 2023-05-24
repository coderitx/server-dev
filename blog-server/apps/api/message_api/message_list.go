package message_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type Message struct {
	SendUserID       uint   `json:"send_user_id"` // 发送人id
	SendUserNickName string `json:"send_user_nick_name"`
	SendUserAvatar   string `json:"send_user_avatar"`

	RevUserID       uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName string    `json:"rev_user_nick_name"`
	RevUserAvatar   string    `json:"rev_user_avatar"`
	Content         string    `json:"content"`    // 消息内容
	CreatedAt       time.Time `json:"created_at"` // 创建时间
	MessageCount    int       `json:"message_count"`
}
type MessageGroup map[uint]*Message

// MessageListView 消息列表
// @Tags 消息管理
// @Summary 消息列表
// @Description 消息列表
// @Router /api/messages [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []Message
	global.DB.Order("created_at asc").Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			// 不存在
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}

	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	responsex.OkWithData(messages, c)
	return
}
