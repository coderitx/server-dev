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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	code := l.svcCtx.RDB.Get(l.ctx, req.Email).Val()
	if code != req.Code {
		logx.Error("the verification code entered is incorrect")
		return nil, err
	}
	nameExists := []string{}
	err = l.svcCtx.DB.Model(models.UserBasic{}).Where("name = ?", req.Name).Pluck("name", &nameExists).Debug().Error
	if err != nil {
		logx.Errorf("query userinfo by name error: %v", err)
		return nil, err
	}
	if len(nameExists) != 0 {
		logx.Error("username already exists")
		return nil, errors.New("username already exists")
	}
	// 所有验证没有问题，则数据入库
	userinfo := models.UserBasic{
		Identity: utils.UUID(),
		Name:     req.Name,
		Password: utils.MD5(req.Password),
		Email:    req.Email,
	}
	err = l.svcCtx.DB.Table(models.UserBasic{}.TableName()).Create(&userinfo).Debug().Error
	if err != nil {
		logx.Errorf("save userinfo to database error: %v", err)
		return nil, err
	}
	return &types.UserRegisterReply{
		Code: constant.SuccessCode,
	}, nil
}
