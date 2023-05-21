package common_svc

import "blog-server/apps/models"

func PageInfoValid(page models.PageInfo) models.PageInfo {
	if page.Page == 0 {
		page.Page = 1
	}
	if page.Limit == 0 {
		page.Limit = 10
	}
	return page
}
