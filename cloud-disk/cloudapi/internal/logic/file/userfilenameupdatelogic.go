package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"context"
	"errors"
	"path"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	var count int
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository AS ur WHERE ur.identity = ?)", req.Name, req.Identity).Count(&count).Debug().Error
	if err != nil {
		logx.Errorf("check if the file name exists error: %v", err)
		return nil, err
	}
	if count > 0 {
		logx.Errorf("the  file name already exists")
		return nil, errors.New("the  file name already exists")
	}
	updatesInfo := models.UserRepository{
		Name: req.Name,
		Ext:  path.Ext(req.Name),
	}
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Updates(updatesInfo).Debug().Error
	if err != nil {
		logx.Errorf("update filename error: %v", err)
		return nil, err
	}
	return
}
