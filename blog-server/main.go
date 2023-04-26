package main

import (
	"blog-server/cmd"
	"blog-server/common/logx"
	"blog-server/config"
	"blog-server/global"
	"blog-server/internal"
	"blog-server/routers"
	"flag"
	"go.uber.org/zap"
)

func init() {
	fp := flag.String("f", "./config/config.yaml", "config file path")
	c, err := config.LoadConfig(*fp)
	if err != nil {
		panic(err)
	}
	global.GlobalC = c
	internal.InitDB(global.GlobalC.Mysql)
	internal.InitRedis(global.GlobalC.Redis)
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

func main() {
	routers.InitApiRouter()
}
