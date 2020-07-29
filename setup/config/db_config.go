package config

import "github.com/spf13/viper"

type DBConfig struct {
	Driver string
	DSN    string
}

func NewDBConfig() *DBConfig {
	db := &DBConfig{
		Driver: "sqlite3",
		DSN:    "test.db",
	}
	_ = viper.UnmarshalKey("db", db)
	return db
}
