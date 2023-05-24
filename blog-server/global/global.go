package global

import (
	"blog-server/config"
	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	ConfigPath = "./config/config.yaml"
	GlobalC    *config.Config
	DB         *gorm.DB
	RDB        *redis.Client
	AddrDB     *geoip2.DBReader
)
