package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LogConfig struct {
	LogLevel          string `yaml:"level" json:"level"`                       // 日志打印级别 debug  info  warning  error
	LogFormat         string `yaml:"format" json:"format"`                     // 输出日志格式	logfmt, json
	LogPath           string `yaml:"path" json:"path"`                         // 输出日志文件路径
	LogFileName       string `yaml:"filename" json:"filename"`                 // 输出日志文件名称
	LogFileMaxSize    int    `yaml:"file_maxsize" json:"file_maxsize"`         // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    `yaml:"file_max_backups" json:"file_max_backups"` // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    `yaml:"max_age" json:"max_age"`                   // 日志保留时间，单位: 天 (day)
	LogCompress       bool   `yaml:"compress" json:"compress"`                 // 是否压缩日志
	LogStdout         bool   `yaml:"stdout" json:"stdout"`                     // 是否输出到控制台
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	LogLevel string `yaml:"log_level"`
}

type SystemConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

type RedisConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Password    string `yaml:"password"`
	ReadTimeout int    `yaml:"read_timeout"`
	DialTimeout int    `yaml:"dial_timeout"`
	PoolSize    int    `yaml:"pool_size"`
	PoolTimeout int    `yaml:"pool_timeout"`
	MaxConnAge  int    `yaml:"max_conn_age"`
}

type Config struct {
	Mysql  MysqlConfig  `yaml:"mysql"`
	System SystemConfig `yaml:"system"`
	Log    LogConfig    `yaml:"log"`
	Redis  RedisConfig  `yaml:"redis"`
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
