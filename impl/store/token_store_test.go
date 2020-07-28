package store

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

var (
	ac  = uuid.New().String()
	at  = uuid.New().String()
	rt  = uuid.New().String()
	req = uuid.New().String()
)

func TestDBStore_StoreTokenProfile(t *testing.T) {
	open, err2 := gorm.Open("sqlite3", "test.db")
	if err2 != nil {
		t.Error(err2)
		return
	}
	defer open.Close()
	store := NewTokenStore(open)
	store.AutoMigrate()
	sign := oidcsdk.TokenSignatures{
		AuthorizationCodeSignature: ac,
		AccessTokenSignature:       at,
		RefreshTokenSignature:      rt,
		RefreshTokenExpiry:         time.Now(),
		AccessTokenExpiry:          time.Now(),
		AuthorizationCodeExpiry:    time.Now(),
	}
	prof := oidcsdk.RequestProfile{}
	prof.SetNonce("jhjhjhjh")
	err := store.StoreTokenProfile(nil, req, sign, prof)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDBStore_GetProfileWithAuthCodeSign(t *testing.T) {
	open, err2 := gorm.Open("sqlite3", "test.db")
	if err2 != nil {
		t.Error(err2)
		return
	}
	defer open.Close()
	store := NewTokenStore(open)
	store.AutoMigrate()

	profile, id, err := store.GetProfileWithAuthCodeSign(nil, ac)
	if err != nil {
		t.Error(err)
	}
	if id == "" {
		t.Error("not found")
	}

	fmt.Println("profile:", profile)
}
