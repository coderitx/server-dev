package config

import (
	"blog-server/config/internal_config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql    internal_config.MysqlConfig  `yaml:"mysql"`
	System   internal_config.SystemConfig `yaml:"system"`
	Log      internal_config.LogConfig    `yaml:"log"`
	Redis    internal_config.RedisConfig  `yaml:"redis"`
	SiteInfo internal_config.SiteInfo     `yaml:"site_info"`
	QQ       internal_config.QQ           `yaml:"qq"`
	Tencent  internal_config.COS          `yaml:"cos"`
	Email    internal_config.Email        `yaml:"email"`
}

func LoadConfig(path string) (*Config, error) {
	C := new(Config)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, C)
	if err != nil {
		fmt.Printf("unmarshal config error: %v \n", err.Error())
		return nil, err
	}
	fmt.Printf("%+v", C)
	return C, nil
}
