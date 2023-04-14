package main

import (
	"online-chat/common/logx"
	"online-chat/config"
	"online-chat/internal"
	"online-chat/online_server/router"
)

func init() {
	logx.InitLogger(config.LogConfigs{})
	internal.InitDB()
	internal.InitRedis()
}

func main() {
	router.InitApiRouter()
}
