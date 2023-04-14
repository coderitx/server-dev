package config

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
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

type RedisConfig struct {
	Addr       string `yaml:"addr" json:"addr"`
	Password   string `yaml:"password" json:"password"`
	MaxConnAge int    `yaml:"max_conn_age" json:"max_conn_age"`
}

type DBConfig struct {
	Addr        string `yaml:"addr" json:"addr"`
	Port        int    `yaml:"port" json:"port"`
	Username    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
	Database    string `yaml:"database" json:"database"`
	MaxConnect  int    `yaml:"max_connect" json:"max_connect"`
	OpenConnect int    `yaml:"open_connect" json:"open_connect"`
}

type Web struct {
	Port int `yaml:"port" json:"port"`
}

type GlobalConfig struct {
	Web      Web         `yaml:"web" json:"web"`
	LogCfg   LogConfig   `yaml:"log" json:"log"`
	DBCfg    DBConfig    `yaml:"db"json:"db"`
	RedisCfg RedisConfig `yaml:"redis"json:"redis"`
}

func LoadConfig(path string) (*GlobalConfig, error) {
	c := new(GlobalConfig)
	file, _ := os.Open(path)
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("-----LoadConfig-----%+v\n", c)
	return c, nil
}
