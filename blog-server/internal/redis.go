package internal

import (
	"blog-server/config"
	"blog-server/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func InitRedis(c config.RedisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        c.Host,
		Password:    c.Password,
		DialTimeout: time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout: time.Duration(c.ReadTimeout) * time.Second,
		PoolSize:    c.PoolSize,
		PoolTimeout: time.Duration(c.PoolTimeout) * time.Second,
		MaxConnAge:  time.Duration(c.MaxConnAge) * time.Second,
	})

	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		zap.S().Errorf("redis connection error: %v", err)
		return
	}
	fmt.Println("=====redis========")
	global.RDB = rdb
	return
}
