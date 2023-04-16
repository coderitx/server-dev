package models

import (
	"online-chat/global"
	"online-chat/online_server/models/requestx"
	"online-chat/online_server/models/wsx"
)

// AutoMigrateModels 迁移模型至数据库
func AutoMigrateModels() {
	global.DB.AutoMigrate(&requestx.UserBasic{})
	global.DB.AutoMigrate(&wsx.WSMessage{})
	global.DB.AutoMigrate(&wsx.Contact{})
	global.DB.AutoMigrate(&wsx.GroupBasic{})
}
