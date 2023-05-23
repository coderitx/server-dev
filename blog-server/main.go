package main

import (
	"blog-server/apps/routers"
	"blog-server/cmd"
	"blog-server/common/logx"
	"blog-server/config"
	_ "blog-server/docs"
	"blog-server/global"
	"blog-server/internal"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	//fp := flag.String("f", "./config/config.yaml", "config file path")
	godotenv.Load("./private.env")
	c, err := config.LoadConfig(global.ConfigPath)
	if err != nil {
		panic(err)
	}
	global.GlobalC = c
	err = internal.InitDB(global.GlobalC.Mysql)
	if err != nil {
		zap.S().Errorf("mysql 初始化失败 [ERROR]: %v", err.Error())
		panic(err)
	}
	err = internal.InitRedis(global.GlobalC.Redis)
	if err != nil {
		zap.S().Errorf("redis 初始化失败 [ERROR]: %v", err.Error())
		panic(err)
	}
	logx.InitLogger(global.GlobalC.Log)
	zap.S().Infoln("--------所有配置初始化完成---------")
	// 迁移数据表
	opt := cmd.Parse()
	// 迁移完成是否停止服务
	if cmd.IsWebStop(opt) {
		cmd.SwitchOption(opt)
		return
	}

}

// @title blog-server API 文档
// @version 1.0
// @description blog-server API 文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	routers.InitApiRouter()
}
