package config

type AppConfig struct {
	Boards    map[string]BoardConfig `mapstructure:"boards"`
	Recaptcha RecaptchaConfig
}

// BoardConfig contains storage connection configuration for relational storage
type BoardConfig struct {
	Name        string `mapstructure:"name"`
	Shorthand   string `mapstructure:"shorthand"`
	Description string `mapstructure:"description"`
	RefCounter  int64  `mapstructure:"ref_counter"`
	VipOnly     bool   `mapstructure:"vip_only"`
}
