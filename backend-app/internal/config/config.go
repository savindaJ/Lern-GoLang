package config

import "os"

type Config struct {
	AppName string
	AppEnv  string
	AppPort string
}

func Load() *Config {
	return &Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
