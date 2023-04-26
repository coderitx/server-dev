package internal_config

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
