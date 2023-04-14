package internal

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

// InitRedis 初始化redis
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	if err := rdb.Ping().Err(); err != nil {
		zap.S().Errorf("redis connection error: %v", err)
		return nil
	}
	return rdb
}
