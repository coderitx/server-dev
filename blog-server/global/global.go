package global

import (
	"blog-server/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GlobalC *config.Config
	DB      *gorm.DB
	RDB     *redis.Client
)
