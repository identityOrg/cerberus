package store

import (
	"encoding/json"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
	"time"
)

type TokenDBModel struct {
	gorm.Model
	RequestID      string                 `gorm:"column:request_id;unique_index;size:60"`
	ATSignature    string                 `gorm:"column:at_signature;unique_index"`
	ACSignature    string                 `gorm:"column:ac_signature;unique_index"`
	RTSignature    string                 `gorm:"column:rt_signature;unique_index"`
	ACExpiry       time.Time              `gorm:"column:ac_expiry"`
	ATExpiry       time.Time              `gorm:"column:at_expiry"`
	RTExpiry       time.Time              `gorm:"column:rt_expiry"`
	ProfileData    string                 `gorm:"column:profile;type:lob"`
	RequestProfile oidcsdk.RequestProfile `gorm:"-"`
}

func (t *TokenDBModel) BeforeCreate() error {
	marshal, err := json.Marshal(t.RequestProfile)
	if err != nil {
		return err
	}
	t.ProfileData = string(marshal)
	return nil
}

func (t *TokenDBModel) AfterFind() (err error) {
	if t.ProfileData != "" {
		var rp oidcsdk.RequestProfile
		err = json.Unmarshal([]byte(t.ProfileData), &rp)
		if err != nil {
			return err
		}
		t.RequestProfile = rp
	}
	return
}

func (t *TokenDBModel) TableName() string {
	return "tokens"
}
