package config

// ServerConfig contains configuration for REST API server
type ServerConfig struct {
	Hostname string         `mapstructure:"hostname"`
	Port     int            `mapstructure:"port"`
	Prefix   string         `mapstructure:"prefix"`
	Database DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig contains storage connection configuration for relational storage
type DatabaseConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	TimeZone    string `mapstructure:"timezone"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}
