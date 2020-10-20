package setup

import (
	"fmt"
	"github.com/identityOrg/cerberus/setup/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB(config *config.DBConfig, serverConfig *config.ServerConfig) (*gorm.DB, error) {
	var orm *gorm.DB
	var err error
	switch config.Driver {
	case "sqlite3":
		orm, err = gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	case "mysql":
		orm, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	case "postgres":
		orm, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	default:
		err = fmt.Errorf("unknown sql dialect/driver %s", config.Driver)
	}
	if err != nil {
		return nil, err
	}
	if serverConfig.Debug {
		orm = orm.Debug()
	}
	return orm, nil
}
