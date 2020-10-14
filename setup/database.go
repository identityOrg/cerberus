package setup

import (
	"fmt"
	"github.com/identityOrg/cerberus/setup/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB(config *config.DBConfig) (*gorm.DB, error) {
	switch config.Driver {
	case "sqlite3":
		return gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	case "mysql":
		return gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	case "postgres":
		return gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unknown sql dialect/driver %s", config.Driver)
	}
}
