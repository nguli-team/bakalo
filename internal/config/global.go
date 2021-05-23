package config

import "sync"

var (
	cfg     Config
	cfgOnce sync.Once
)

func SetGlobalConfig(newConfig Config) {
	cfgOnce.Do(func() {
		cfg = newConfig
	})
}

func Env() Environment {
	return cfg.Env
}

func Server() HTTPServerConfig {
	return cfg.Server
}

func Database() DatabaseConfig {
	return cfg.Database
}

func App() AppConfig {
	return cfg.App
}
