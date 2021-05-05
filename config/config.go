package config

import (
	"github.com/spf13/viper"
)

// Config contains cmd configurations such as server, database, etc.
type Config struct {
	Env    Environment  `mapstructure:"environment"`
	Server ServerConfig `mapstructure:"server"`
	App    AppConfig    `mapstructure:"cmd"`
}

// NewConfig creates new configuration struct from YAML file.
func NewConfig(filename string) (*Config, error) {
	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.checkConfig()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// checkConfig checks configurations validity and provides defaults if necessary.
func (c *Config) checkConfig() error {
	err := c.Env.IsValid()
	if err != nil {
		return err
	}
	// TODO(core): (Okka) actually implements config checks
	return nil
}
