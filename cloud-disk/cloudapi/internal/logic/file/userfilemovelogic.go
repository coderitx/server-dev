package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	/*
		1. 查询目标目录和西东文件的的信息
		2. 判断目标目录下下是否已经存在即将移动的文件的名称
		3. 存在则直接返回，不存在则更新需要移动的文件的parent_id 为父级目录的parent_id
	*/

	var parentUserRepository []models.UserRepository
	var subUserRepository []models.UserRepository
	// 查询目标目录的信息
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ? AND user_identity = ?", req.ParentIdnetity, userIdentity).Scan(&parentUserRepository).Debug().Error
	// 需要移动的目录的信息
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ? AND user_identity = ?", req.Idnetity, userIdentity).Scan(&subUserRepository).Debug().Error
	if err != nil {
		return nil, err
	}
	// 如果父级目录不存在
	if len(parentUserRepository) == 0 {
		return nil, errors.New("the file not found")
	}
	// 如果需要移动的目录不存在
	if len(subUserRepository) == 0 {
		return nil, errors.New("move file not found")
	}

	ur := parentUserRepository[0]
	var count int
	// 判断目标路径下是否已经存在即将移动的目录或者文件
	err = l.svcCtx.DB.Model(models.UserRepository{}).Where("identity = ? AND user_identity = ? AND parent_id = ? AND name = ?", req.ParentIdnetity, userIdentity, ur.ParentId, subUserRepository[0].Name).Count(&count).Debug().Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("the  file name already exists")
	}
	// 更新需要移动的文件或者目录的信息
	newUR := models.UserRepository{
		ParentId: int64(ur.ID),
	}
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Where("identity = ? AND user_identity = ?", req.Idnetity, userIdentity).Update(&newUR).Debug().Error
	if err != nil {
		return nil, err
	}
	return
}
