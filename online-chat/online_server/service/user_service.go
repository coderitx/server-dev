package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"online-chat/common/errorx"
	"online-chat/common/response"
	"online-chat/online_server/models/requestx"
	"online-chat/online_server/models/responsex"
	"online-chat/utils"
)

// GetUserList
// @Success 200 {string} json{"code","message","data"}
// @Router /user/getUserList [get]
func GetUserList(ctx *gin.Context) {
	u := requestx.UserBasic{}
	rsp, code := u.GetUserList()
	response.Success(ctx, http.StatusOK, errorx.ErrMsg(code), rsp)
	return
}

// CreateUser
// @Success 200 {string} json{"code","message","data"}
// @Router /user/createUser [post]
func CreateUser(ctx *gin.Context) {
	u := requestx.UserBasic{}

	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	code := u.CreateUser()
	response.Success(ctx, code, errorx.ErrMsg(code), nil)
	return
}

// DeleteUser
// @Success 200 {string} json{"code","message","data"}
// @Router /user/deleteUser [delete]
func DeleteUser(ctx *gin.Context) {
	u := requestx.UserBasic{}

	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	code := u.DeleteUser()
	response.Success(ctx, code, errorx.ErrMsg(code), nil)
	return
}

// UpdateUser
// @Success 200 {string} json{"code","message","data"}
// @Router /user/deleteUser [put]
func UpdateUser(ctx *gin.Context) {
	u := requestx.UserBasic{}

	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	code := u.UpdateUser()
	response.Success(ctx, code, errorx.ErrMsg(code), nil)
	return
}

// GetUserByName
// @Success 200 {string} json{"code","message","data"}
// @Router /user/getUser.name [get]
func GetUserByName(ctx *gin.Context) {
	u := requestx.UserBasic{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	user, code := u.FindUserByName()
	response.Success(ctx, code, errorx.ErrMsg(code), user)
	return
}

// GetUserByID
// @Success 200 {string} json{"code","message","data"}
// @Router /user/getUser.id [get]
func GetUserByID(ctx *gin.Context) {
	u := requestx.UserBasic{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	user, code := u.FindUserByID()
	response.Success(ctx, code, errorx.ErrMsg(code), user)
	return
}

// GetUserByPhone
// @Success 200 {string} json{"code","message","data"}
// @Router /user/getUser.phone [get]
func GetUserByPhone(ctx *gin.Context) {
	u := requestx.UserBasic{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	user, code := u.FindUserByPhone()
	response.Success(ctx, code, errorx.ErrMsg(code), user)
	return
}

// GetUserByEmail
// @Success 200 {string} json{"code","message","data"}
// @Router /user/getUser.email [get]
func GetUserByEmail(ctx *gin.Context) {
	u := requestx.UserBasic{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	user, code := u.FindUserByEmail()
	response.Success(ctx, code, errorx.ErrMsg(code), user)
	return
}

// Login
// @Success 200 {string} json{"code","message","data"}
// @Router /user/login [post]
func Login(ctx *gin.Context) {
	u := requestx.UserBasic{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		zap.S().Errorf("bind request data error: %v", err)
		response.Failed(ctx, http.StatusOK, errorx.ParameterErrorCode, errorx.ErrMsg(errorx.ParameterErrorCode), nil)
		return
	}
	user, code := u.FindUserByNameOrEmailOrPhoneAndPwd()
	if code != errorx.SuccessCode {
		response.Failed(ctx, http.StatusOK, code, errorx.ErrMsg(code), nil)
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Identity, user.Name, user.Phone, user.Email, 300)
	if err != nil {
		response.Failed(ctx, http.StatusOK, errorx.ServerErrorCode, errorx.ErrMsg(errorx.ServerErrorCode), nil)
		return
	}
	reply := responsex.LoginReply{
		Name:  user.Name,
		Phone: user.Phone,
		Token: token,
	}
	response.Success(ctx, http.StatusOK, errorx.ErrMsg(errorx.SuccessCode), reply)
}
