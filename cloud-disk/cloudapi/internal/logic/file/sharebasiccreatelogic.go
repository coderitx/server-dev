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

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	// 根据user_repository 的identity信息创建了一条分享记录
	var ur []models.UserRepository
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ?", req.UserRepositoryIdentity).Scan(&ur).Debug().Error
	if err != nil {
		return nil, err
	}
	if len(ur) == 0 {
		return nil, errors.New("repository not found")
	}
	uuid := utils.UUID()
	shareBasicData := models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		RepositoryIdentity:     ur[0].RepositoryIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	err = l.svcCtx.DB.Model(models.ShareBasic{}).Create(&shareBasicData).Debug().Error
	if err != nil {
		return nil, err
	}
	return &types.ShareBasicCreateReply{
		Identity: uuid,
	}, nil
}
