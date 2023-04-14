package file

import (
	"cloud-disk/cloudapi/constant"
	"cloud-disk/cloudapi/utils"
	"context"

	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationReply, err error) {
	u, err := utils.AnalyzeToken(authorization)
	if err != nil {
		return nil, err
	}
	token, _ := utils.GenerateToken(u.Id, u.Identity, u.Name, constant.CodeExpire)
	refreshToken, _ := utils.GenerateToken(u.Id, u.Identity, u.Name, constant.CodeExpire)
	return &types.RefreshAuthorizationReply{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
