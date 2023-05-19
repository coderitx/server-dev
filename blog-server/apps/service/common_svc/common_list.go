package common_svc

import (
	"blog-server/apps/models"
	"blog-server/global"
)

type Options struct {
	models.PageInfo
}

func ComList[T any](obj T, opt Options) (list []T, count int64, err error) {
	count = global.DB.Select("id").Find(&obj).RowsAffected
	page := (opt.Page - 1) * opt.Limit
	if page < 0 {
		page = 0
	}
	if opt.Sort == "" {
		opt.Sort = "created_at desc"
	}
	err = global.DB.Debug().Limit(opt.Limit).Offset(page).Order(opt.Sort).Find(&list).Error
	return list, count, err
}
