package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"cloud-disk/cloudapi/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	repo := models.RepositoryPool{
		Name:     req.Name,
		Size:     req.Size,
		Ext:      req.Ext,
		Hash:     req.Hash,
		Identity: utils.UUID(),
		Path:     req.Path,
	}
	err = l.svcCtx.DB.Model(models.RepositoryPool{}).Create(&repo).Debug().Error
	if err != nil {
		logx.Errorf("insert fileinfo to database error: %v", err)
		return nil, err
	}
	return &types.FileUploadReply{
		Identity: repo.Identity,
		Name:     repo.Name,
		Ext:      repo.Ext,
	}, nil
}
