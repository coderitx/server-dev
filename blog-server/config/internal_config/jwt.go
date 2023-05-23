package internal_config

type JWT struct {
	Secret  string `yaml:"secret" json:"secret"`
	Expires int64  `yaml:"expires" json:"expires"`
	IsUser  string `yaml:"is_user" json:"is_user"`
}
