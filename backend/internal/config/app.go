package config

import "github.com/spf13/viper"

type AppConfig struct {
	Name   string
	Port   string
	Env    string
	Secret string
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Name:   viper.GetString("app.name"),
		Port:   viper.GetString("app.port"),
		Env:    viper.GetString("app.env"),
		Secret: viper.GetString("app.secret"),
	}
}
