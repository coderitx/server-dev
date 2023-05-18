package cmd

import (
	models2 "blog-server/apps/models"
	"blog-server/global"
	"go.uber.org/zap"
)

// Makemigrations 迁移表
func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models2.UserModel{}, "CollectsModels", &models2.UserCollectModel{})
	global.DB.SetupJoinTable(&models2.MenuModel{}, "Banners", &models2.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models2.BannerModel{},
			&models2.TagModel{},
			&models2.MessageModel{},
			&models2.AdvertModel{},
			&models2.UserModel{},
			&models2.CommentModel{},
			&models2.ArticleModel{},
			&models2.MenuModel{},
			&models2.MenuBannerModel{},
			&models2.FadeBackModel{},
			&models2.LoginDataModel{},
		)
	if err != nil {
		zap.S().Error("[ error ] 生成数据库表结构失败！")
		return
	}
	zap.S().Infof("[ success ] 生成数据库表结构成功！")
}
