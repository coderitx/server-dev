package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"cloud-disk/cloudapi/utils"
	"context"

	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	var count int
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(&count).Debug().Error
	if err != nil {
		logx.Errorf("check if the file name exists error: %v", err)
		return nil, err
	}
	if count > 0 {
		logx.Errorf("the  file name already exists")
		return nil, errors.New("the  file name already exists")
	}
	repository := models.UserRepository{
		Identity:     utils.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	err = l.svcCtx.DB.Model(models.UserRepository{}).Create(&repository).Debug().Error
	if err != nil {
		logx.Errorf("create name = %v folder error: %v", req.Name, err)
		return nil, err
	}
	return
}
