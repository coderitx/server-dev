package file

import (
	"cloud-disk/cloudapi/constant"
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	var userinfo []string
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("user_identity = ?", userIdentity).Pluck("parent_id", &userinfo).Debug().Error
	if err != nil {
		logx.Errorf("query user id error: %v", err)
		return nil, err
	}
	if len(userinfo) == 0 {
		logx.Errorf("username not found.")
		return nil, errors.New("username not found")
	}
	page := req.Page
	if page == 0 {
		page = constant.Page
	}
	size := req.Size
	if size == 0 {
		size = constant.Size
	}
	offset := (page - 1) * size
	var uf []*types.UserFile
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("parent_id = ? AND user_identity = ? ", userinfo[0], userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,user_repository.name, repository_pool.path, repository_pool.size").
		Joins("LEFT JOIN repository_pool ON user_repository.repository_identity=repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Now().Local()).
		Limit(size).Offset(offset).
		Scan(&uf).Debug().Error
	if err != nil {
		logx.Errorf("query user file list error: %v", err)
		return nil, err
	}
	var count int64
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("parent_id = ? AND user_identity = ? ", userinfo[0], userIdentity).Count(&count).Debug().Error
	if err != nil {
		logx.Errorf("query user file list count error: %v", err)
		return nil, err
	}
	rsp := &types.UserFileListReply{
		List:  uf,
		Count: count,
	}
	return rsp, nil
}
