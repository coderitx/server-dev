package user_api

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

const Avatar = "/uploads/avatar/default.jpg"

func (*UserApi) UserCreateView(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responsex.FailWithError(err, &req, c)
		return
	}
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", req.UserName).Error
	if err == nil {
		responsex.FailWithMessage("用户民已存在", c)
		return
	}
	passMd5 := utils.MD5([]byte(req.Password))
	err = global.DB.Create(&models.UserModel{
		NickName:   req.NickName,
		UserName:   req.UserName,
		Password:   passMd5,
		Email:      "",
		Role:       req.Role,
		Avatar:     Avatar,
		IP:         c.ClientIP(),
		Addr:       utils.GetAddr(c.ClientIP()),
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		responsex.FailWithMessage("创建用户失败", c)
		return
	}
	responsex.OkWithData("创建用户成功", c)
	return
}
