package config

type Command struct {
	LogoutPath string `mapstructure:"logout-path" json:"logoutPath" yaml:"logout-path"`
}
