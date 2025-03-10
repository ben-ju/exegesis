package config

import "os"

type Config struct {
	AppURL     string
	AppPort    string
	PsqlHost   string
	PsqlPort   string
	PsqlUser   string
	PsqlPass   string
	PsqlDBName string
	LogPath    string
	AppKey     string
}

func NewConfig() Config {
	return Config{
		AppURL:     os.Getenv("APP_URL"),
		AppPort:    os.Getenv("APP_PORT"),
		PsqlHost:   os.Getenv("POSTGRES_HOST"),
		PsqlPort:   os.Getenv("POSTGRES_PORT"),
		PsqlUser:   os.Getenv("POSTGRES_USER"),
		PsqlPass:   os.Getenv("POSTGRES_PASSWORD"),
		PsqlDBName: os.Getenv("POSTGRES_DB"),
		LogPath:    os.Getenv("LOG_PATH"),
		AppKey:     os.Getenv("APP_KEY"),
	}
}
