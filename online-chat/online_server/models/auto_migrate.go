package models

import (
	"online-chat/config"
)

// AutoMigrateModels 迁移模型至数据库
func AutoMigrateModels() {
	gormDB := config.InitDB()
	gormDB.AutoMigrate(&UserBasic{})
}
