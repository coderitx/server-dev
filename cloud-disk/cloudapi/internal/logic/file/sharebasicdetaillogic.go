package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	// 更新点击次数
	err = l.svcCtx.DB.Model(models.ShareBasic{}).Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity).Debug().Error
	if err != nil {
		return nil, err
	}
	// 查询分享的资源的详细信息
	// 存在重命名问题，所以name使用user_repository表中的名称
	var shareBasicDetails []types.ShareBasicDetailReply
	err = l.svcCtx.DB.Model(models.ShareBasic{}).Where("share_basic.identity = ?", req.Identity).
		Select("share_basic.repository_identity, repository_pool.ext,user_repository.name, repository_pool.path, repository_pool.size").
		Joins("LEFT JOIN repository_pool ON repository_pool.identity=share_basic.repository_identity").
		Joins("LEFT JOIN user_repository ON share_basic.user_repository_identity=user_repository.identity").
		Where("share_basic.deleted_at = ? OR share_basic.deleted_at IS NULL", time.Now().Local()).
		Scan(&shareBasicDetails).Debug().Error
	if err != nil {
		return nil, err
	}
	if len(shareBasicDetails) == 0 {
		return nil, errors.New("share file not found")
	}

	return &shareBasicDetails[0], nil
}
