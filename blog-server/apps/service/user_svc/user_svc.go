package user_svc

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/global"
	"blog-server/utils"
	"fmt"
	"go.uber.org/zap"
)

type UserService struct {
}

func (*UserService) CreateUser(userName, nickName, password string, role ctype.Role, email, host string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		zap.S().Errorf("%v 用户名已存在", userName)
		return fmt.Errorf("%v 用户名已存在", userName)
	}
	password = utils.MD5([]byte(password))
	err = global.DB.Create(&models.UserModel{
		UserName: userName,
		NickName: nickName,
		Password: password,
		Role:     role,
		Email:    email,
		Addr:     host,
	}).Error
	if err != nil {
		zap.S().Error("创建用户失败")
		return err
	}
	return nil
}
