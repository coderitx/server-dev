package file

import (
	"cloud-disk/cloudapi/internal/models"
	"context"
	"errors"

	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteReply, err error) {
	var count int
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Count(&count).Debug().Error
	if err != nil {
		logx.Errorf("check if the file name exists error: %v", err)
		return nil, err
	}
	if count == 0 {
		logx.Errorf("no such file or directory")
		return nil, errors.New("no such file or directory")
	}
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Delete(&models.UserRepository{}).Debug().Error
	if err != nil {
		logx.Errorf("delete file error: %v", err)
		return nil, err
	}
	return
}
