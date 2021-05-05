package config

type AppConfig struct {
	Boards map[string]BoardConfig `mapstructure:"boards"`
}

// BoardConfig contains database connection configuration for relational database
type BoardConfig struct {
	Title       string `mapstructure:"title"`
	Short       string `mapstructure:"short"`
	Description string `mapstructure:"description"`
	VipOnly     bool   `mapstructure:"vip_only"`
}
