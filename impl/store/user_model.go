package store

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserDBModel struct {
	gorm.Model
	Username          string     `gorm:"column:username;unique_index;size:256;not null"`
	Password          string     `gorm:"column:password;size:256;not null"`
	Locked            bool       `gorm:"column:locked"`
	Blocked           bool       `gorm:"column:blocked"`
	WrongAttemptStart *time.Time `gorm:"column:wrong_attempt_start"`
	WrongAttemptCount uint8      `gorm:"column:wrong_attempt_count"`
}

func (t *UserDBModel) TableName() string {
	return "users"
}
