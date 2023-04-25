package main

import (
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
	flag.Parse()
	c, err := config.LoadConfig(*fp)
	if err != nil {
		panic(err)
	}
	global.GlobalC = c
	internal.InitDB(global.GlobalC.Mysql)
	internal.InitRedis(global.GlobalC.Redis)
	logx.InitLogger(global.GlobalC.Log)
	zap.S().Infoln("--------所有配置初始化完成---------")
}

func main() {
	routers.InitApiRouter()
}
