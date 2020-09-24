package setup

import (
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewGormDB(config *config.DBConfig) (*gorm.DB, error) {
	return gorm.Open(config.Driver, config.DSN)
}
