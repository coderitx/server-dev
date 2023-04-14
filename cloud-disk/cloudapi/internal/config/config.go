package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Database struct {
		Source string
	}
	RedisConfig struct {
		Addr     string
		Password string
		DB       int
	}
}
