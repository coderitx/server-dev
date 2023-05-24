package message_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`      // 消息内容
}

// MessageCreateView 发布消息
func (*MessageApi) MessageCreateView(c *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登陆人ID
	var req MessageRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	var sendUser, recvUser models.UserModel
	err = global.DB.Take(&sendUser, req.SendUserID).Error
	if err != nil {
		responsex.FailWithMessage("发送人不存在", c)
	}
	err = global.DB.Take(&recvUser, req.RevUserID).Error
	if err != nil {
		responsex.FailWithMessage("接收人不存在", c)
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       req.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        req.RevUserID,
		RevUserNickName:  recvUser.NickName,
		RevUserAvatar:    recvUser.Avatar,
		IsRead:           false,
		Content:          req.Content,
	}).Error
	if err != nil {
		zap.S().Error(err)
		responsex.FailWithMessage("消息发送失败", c)
		return
	}
	responsex.OkWithMessage("消息发送成功", c)
	return
}
