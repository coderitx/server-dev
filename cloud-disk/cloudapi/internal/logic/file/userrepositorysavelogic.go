package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"cloud-disk/cloudapi/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	userRepository := models.UserRepository{
		Name:               req.Name,
		Identity:           utils.UUID(),
		UserIdentity:       userIdentity,
		Ext:                req.Ext,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
	}
	err = l.svcCtx.DB.Model(models.UserRepository{}).Create(&userRepository).Debug().Error
	return
}
