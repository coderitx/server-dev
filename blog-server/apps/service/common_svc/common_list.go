package common_svc

import (
	"blog-server/apps/models"
	"blog-server/global"
)

type Options struct {
	models.PageInfo
}

func ComList[T any](obj T, opt Options) (list []T, count int64, err error) {
	count = global.DB.Select("id").Find(&list).Debug().RowsAffected
	limit := opt.Limit
	offset := (opt.Page - 1) * limit
	if opt.Sort == "" {
		opt.Sort = "created_at desc"
	}
	query := global.DB.Where(obj)
	err = query.Debug().Limit(limit).Offset(offset).Order(opt.Sort).Find(&list).Error
	return list, count, err
}
