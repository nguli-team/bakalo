package config

type SMTPConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	SenderName string `mapstructure:"sender_name"`
	Email      string `mapstructure:"email"`
	Password   string `mapstructure:"password"`
}
