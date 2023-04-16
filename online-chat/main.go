package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"online-chat/common/logx"
	"online-chat/config"
	"online-chat/global"
	"online-chat/internal"
	"online-chat/mock"
	"online-chat/online_server/models"
	"online-chat/online_server/router"
)

func main() {
	configFile := flag.String("f", "config/config.json", "the config file")
	flag.Parse()

	c, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println("解析配置失败")
		panic(err)
	}
	global.GlobalC = c
	logx.InitLogger(global.GlobalC.LogCfg)
	internal.InitDB(global.GlobalC.DBCfg)
	internal.InitRedis(global.GlobalC.RedisCfg)
	zap.S().Info("配置初始化完成")

	models.AutoMigrateModels()
	mock.MockUserBasic()
	router.InitApiRouter()
}
