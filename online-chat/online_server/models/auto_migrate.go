package models

import (
	"online-chat/global"
	"online-chat/online_server/models/requestx"
)

// AutoMigrateModels 迁移模型至数据库
func AutoMigrateModels() {
	global.DB.AutoMigrate(&requestx.UserBasic{})
}
