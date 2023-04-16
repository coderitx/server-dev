package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"online-chat/config"
)

var (
	GlobalC *config.GlobalConfig
	DB      *gorm.DB
	RDB     *redis.Client
)
