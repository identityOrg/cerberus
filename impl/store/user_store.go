package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserStore struct {
	db                 *gorm.DB
	WrongAttemptWindow time.Duration
	WrongAttemptCount  uint8
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db, WrongAttemptWindow: 2 * time.Hour, WrongAttemptCount: 5}
}

func (u *UserStore) AutoMigrate() {
	u.db.AutoMigrate(UserDBModel{})
}

func (u *UserStore) Authenticate(_ context.Context, username string, credential []byte) (err error) {
	user := &UserDBModel{
		Username: username,
	}

	userCondition := u.db.Where(user)
	err = userCondition.Find(user).Error
	if err != nil {
		return err
	}
	if user.ID > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), credential)
		if err != nil {
			now := time.Now()
			if user.WrongAttemptStart == nil || user.WrongAttemptStart.Add(u.WrongAttemptWindow).Before(now) {
				user.WrongAttemptStart = &now
				user.WrongAttemptCount = 1
			} else {
				user.WrongAttemptCount += 1
			}
			if user.WrongAttemptCount > u.WrongAttemptCount {
				user.Locked = true
				err = errors.New("user is locked due to consecutive failure attempt")
			}
			u.db.Update(user)
			return err
		}
		if user.Blocked {
			return errors.New(fmt.Sprintf("user '%s' is blocked", user.Username))
		}
		if user.Locked {
			return errors.New(fmt.Sprintf("user '%s' is locked", user.Username))
		}
	} else {
		return errors.New("credential mismatch")
	}
	return nil
}

func (u *UserStore) GetClaims(context.Context, string, oidcsdk.Arguments, []string) (map[string]interface{}, error) {
	claims := make(map[string]interface{})
	return claims, nil
}

func (u *UserStore) IsConsentRequired(context.Context, string, string, oidcsdk.Arguments) bool {
	return false
}

func (u *UserStore) StoreConsent(context.Context, string, string, oidcsdk.Arguments) error {
	return nil
}

func (u *UserStore) FetchUserProfile(_ context.Context, username string) oidcsdk.RequestProfile {
	rp := oidcsdk.RequestProfile{}
	rp.SetUsername(username)
	return rp
}
