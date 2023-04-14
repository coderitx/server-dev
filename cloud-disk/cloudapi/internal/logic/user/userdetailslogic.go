package user

import (
	"cloud-disk/cloudapi/internal/models"
	"context"

	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailsLogic {
	return &UserDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailsLogic) UserDetails(req *types.UserDetailsRequest) (resp *types.UserDetailsReply, err error) {
	userinfo := []models.UserBasic{}
	err = l.svcCtx.DB.Model(models.UserBasic{}).Where("identity = ?", req.Identity).Scan(&userinfo).Debug().Error
	if err != nil {
		logx.Errorf("query userinfo by identity error: %v", err)
		return nil, err
	}
	if len(userinfo) == 0 {
		logx.Infof("query identity=%v not found", req.Identity)
		return nil, err
	}
	u := userinfo[0]
	return &types.UserDetailsReply{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
