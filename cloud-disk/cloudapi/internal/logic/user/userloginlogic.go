package user

import (
	"cloud-disk/cloudapi/constant"
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"cloud-disk/cloudapi/utils"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	userinfo := []models.UserBasic{}
	err = l.svcCtx.DB.Model(models.UserBasic{}).Where("name=? and password=?", req.Name, utils.MD5(req.Password)).Debug().Scan(&userinfo).Error
	if err != nil {
		logx.Errorf("query user info error: %v", err)
		return nil, err
	}
	if len(userinfo) == 0 {
		logx.Infof("用户名或者密码错误")
		return nil, errors.New("用户名或者密码错误")
	}
	u := userinfo[0]
	token, err := utils.GenerateToken(uint64(u.ID), u.Identity, u.Name, constant.CodeExpire)
	return &types.LoginReply{
		Token: token,
	}, nil
}
