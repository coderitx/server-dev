package service

import (
	"github.com/gin-gonic/gin"
	"online-chat/common/errorx"
	"online-chat/common/response"
	"online-chat/online_server/models/requestx"
)

// GetUserList
// @Success 200 {string} json{"code","message","data"}
// @Router /user/get_user_list [get]
func GetUserList(ctx *gin.Context) {
	u := requestx.UserBasic{}
	rsp, err := u.GetUserList()
	if err != nil {
		response.SendError(ctx, 200, errorx.ErrorCode, "error", nil)
	}
	response.SendSuccess(ctx, 200, "success", rsp)
}
