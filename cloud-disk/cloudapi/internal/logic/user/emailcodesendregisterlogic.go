package user

import (
	"cloud-disk/cloudapi/constant"
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/utils"
	"context"
	"time"

	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeSendRegisterLogic {
	return &EmailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeSendRegisterLogic) EmailCodeSendRegister(req *types.EmailCodeSendRequest) (resp *types.EmailCodeSendResply, err error) {
	var emailExists []string
	err = l.svcCtx.DB.Model(models.UserBasic{}).Select([]string{"email"}).Where("email = ?", req.Email).Pluck("email", &emailExists).Debug().Error
	if err != nil {
		logx.Errorf("query email exists by %v error %v", req.Email, err)
		return nil, err
	}
	if len(emailExists) != 0 {
		logx.Errorf("email: %v already registered", req.Email)
		return nil, err
	}
	code := utils.RandCode()
	err = utils.SendCodeToEmail(req.Email, code)
	if err != nil {
		logx.Errorf("send code to %v error: %v", req.Email, err)
		return &types.EmailCodeSendResply{
			Code: constant.ErrorCode,
		}, err
	}
	err = l.svcCtx.RDB.SetNX(l.ctx, req.Email, code, time.Duration(constant.CodeExpire)*time.Second).Err()
	if err != nil {
		logx.Errorf("save code to redis error: %v", err)
	}
	return &types.EmailCodeSendResply{
		Code: constant.SuccessCode,
	}, nil
	return
}
