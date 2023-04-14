package internal

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"online-chat/config"
	"online-chat/global"
	"time"
)

// InitRedis 初始化redis
func InitRedis(c config.RedisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		MaxConnAge:   time.Duration(c.MaxConnAge),
	})

	if err := rdb.Ping().Err(); err != nil {
		zap.S().Errorf("redis connection error: %v", err)
		return
	}
	global.RDB = rdb
	return
}
