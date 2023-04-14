package svc

import (
	"cloud-disk/cloudapi/internal/config"
	"cloud-disk/cloudapi/internal/middleware"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open("mysql", c.Database.Source)
	if err != nil {
		logx.Errorf("connect mysql error: %v", err)
		return nil
	}
	db.LogMode(true)
	redisCli := redis.NewClient(&redis.Options{
		Addr:     c.RedisConfig.Addr,
		Password: c.RedisConfig.Password,
		DB:       c.RedisConfig.DB,
	})
	return &ServiceContext{
		Config: c,
		DB:     db,
		RDB:    redisCli,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
