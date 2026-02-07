package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Username  string
	Password  string
	Host      string
	Port      int
	Name      string
	Charset   string
	ParseTime bool
	Loc       string
	Idle      int
	Max       int
	Lifetime  int
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:      viper.GetString("database.host"),
		Port:      viper.GetInt("database.port"),
		Username:  viper.GetString("database.username"),
		Password:  viper.GetString("database.password"),
		Name:      viper.GetString("database.name"),
		Charset:   viper.GetString("database.charset"),
		ParseTime: viper.GetBool("database.parseTime"),
		Loc:       viper.GetString("database.loc"),
		Idle:      viper.GetInt("database.idle"),
		Max:       viper.GetInt("database.max"),
		Lifetime:  viper.GetInt("database.lifetime"),
	}
}
