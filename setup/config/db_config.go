package config

type DBConfig struct {
	Driver string
	DSN    string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Driver: "sqlite3",
		DSN:    "test.db",
	}
}
