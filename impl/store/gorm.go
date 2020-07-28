package store

import "github.com/jinzhu/gorm"

func NewGOrmDB(dbType string, dsn string) (*gorm.DB, error) {
	return gorm.Open(dbType, dsn)
}
