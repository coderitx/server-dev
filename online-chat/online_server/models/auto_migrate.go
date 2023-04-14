package models

import (
	"online-chat/global"
	"online-chat/online_server/models/request"
)

// AutoMigrateModels 迁移模型至数据库
func AutoMigrateModels() {
	global.DB.AutoMigrate(&request.UserBasic{})
}
