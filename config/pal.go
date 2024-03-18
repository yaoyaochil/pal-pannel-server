package config

type Palu struct {
	// 文件夹路径
	Init       bool   `mapstructure:"init" json:"Init" yaml:"init"`
	ServerPath string `mapstructure:"server-path" json:"ServerPath" yaml:"server-path"`
}
