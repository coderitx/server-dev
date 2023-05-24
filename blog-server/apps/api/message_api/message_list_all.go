package message_api

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/apps/service/common_svc"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MessageListAllView 所有消息列表
// @Tags 消息管理
// @Summary 所有消息列表
// @Description 所有消息列表
// @Router /api/message_all [get]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (*MessageApi) MessageListAllView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims, ok := _claims.(*utils.CustomClaims)
	if !ok {
		responsex.FailWithMessage("登陆信息有误", c)
		return
	}
	var p models.PageInfo
	err := c.ShouldBindQuery(&p)
	if err != nil {
		zap.S().Errorf("shoud bind params error: %v", err.Error())
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	list, count, err := common_svc.ComList(models.MessageModel{}, common_svc.Options{
		PageInfo: common_svc.PageInfoValid(p),
	})
	if err != nil {
		zap.S().Error("query message list error: %v", err.Error())
		responsex.FailWithCode(errorx.SettingsError, c)
		return
	}

	if claims.Role == int(ctype.PermissionAdmin) {
		responsex.OkWithList(list, count, c)
		return
	}
	listData := []models.MessageModel{}
	for _, msg := range list {
		if msg.SendUserID == claims.UserID || msg.RevUserID == claims.UserID {
			listData = append(list, msg)
		}
	}
	responsex.OkWithList(listData, count, c)
	return
}
