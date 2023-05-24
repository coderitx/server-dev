package internal_config

type RedisConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DB          int    `yaml:"db"`
	Password    string `yaml:"password"`
	ReadTimeout int    `yaml:"read_timeout"`
	DialTimeout int    `yaml:"dial_timeout"`
	PoolSize    int    `yaml:"pool_size"`
	PoolTimeout int    `yaml:"pool_timeout"`
	MaxConnAge  int    `yaml:"max_conn_age"`
}
