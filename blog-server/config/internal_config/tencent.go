package internal_config

type COS struct {
	Enable    bool    `json:"enable" yaml:"enable"` // 是否启用腾讯云存储
	SecretID  string  `json:"secret_id" yaml:"secret_id"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` // 存储桶的名字
	CDN       string  `json:"cdn" yaml:"cdn"`       // 访问图片的地址的前缀
	Prefix    string  `json:"prefix" yaml:"prefix"` // 前缀
	Size      float64 `json:"size" yaml:"size"`     // 存储的大小限制，单位是MB
}
