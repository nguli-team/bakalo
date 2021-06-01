package config

import (
	"github.com/spf13/viper"
)

// Config contains cmd configurations such as server, storage, etc.
type Config struct {
	Env      Environment      `mapstructure:"environment"`
	Server   HTTPServerConfig `mapstructure:"server"`
	Database DatabaseConfig   `mapstructure:"database"`
	App      AppConfig        `mapstructure:"app"`
	SMTP     SMTPConfig       `mapstructure:"smtp"`
}

// NewConfig creates new configuration struct from YAML file.
func NewConfig(filename string) (config Config, err error) {
	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	// TODO: move recaptcha config init
	config.App.Recaptcha.Secret = "6LeogPkaAAAAANvMlB0t5KCoYowah14gxQemUU-E"

	err = config.checkConfig()
	if err != nil {
		return
	}

	return
}

// checkConfig checks configurations validity and provides defaults if necessary.
func (c *Config) checkConfig() error {
	err := c.Env.IsValid()
	if err != nil {
		return err
	}
	// TODO(core): (Okka) actually implements Config checks
	return nil
}
