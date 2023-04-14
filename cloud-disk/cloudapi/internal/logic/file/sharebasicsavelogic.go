package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/utils"
	"context"
	"errors"

	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	// 获取资源的详情
	var repositoryInfo []models.RepositoryPool
	err = l.svcCtx.DB.Model(models.RepositoryPool{}).Where("identity = ?", req.RepositoryIdentity).Scan(&repositoryInfo).Debug().Error
	if err != nil {
		return nil, err
	}
	if len(repositoryInfo) == 0 {
		return nil, errors.New("source not found")
	}
	// user_repository 资源保存
	rp := repositoryInfo[0]
	identity := utils.UUID()
	userRepository := models.UserRepository{
		Identity:           identity,
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: rp.Identity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	err = l.svcCtx.DB.Model(models.UserRepository{}).Create(&userRepository).Debug().Error
	if err != nil {
		return nil, err
	}

	return &types.ShareBasicSaveReply{
		Identity: identity,
	}, nil
}
