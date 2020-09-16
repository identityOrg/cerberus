package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
	"time"
)

type TokenStore struct {
	db *gorm.DB
}

func NewTokenStore(db *gorm.DB) *TokenStore {
	return &TokenStore{db: db}
}

func (d *TokenStore) AutoMigrate() {
	d.db.AutoMigrate(TokenDBModel{})
}

func (d *TokenStore) StoreTokenProfile(_ context.Context, reqId string, signatures oidcsdk.ITokenSignatures, profile oidcsdk.RequestProfile) (err error) {
	tm := &TokenDBModel{
		RequestID:      reqId,
		ACSignature:    toNullString(signatures.GetACSignature()),
		ACExpiry:       signatures.GetACExpiry(),
		ATSignature:    toNullString(signatures.GetATSignature()),
		ATExpiry:       signatures.GetATExpiry(),
		RTSignature:    toNullString(signatures.GetRTSignature()),
		RTExpiry:       signatures.GetRTExpiry(),
		RequestProfile: profile,
	}
	return d.db.Create(tm).Error
}

func (d *TokenStore) GetProfileWithAuthCodeSign(_ context.Context, signature string) (profile oidcsdk.RequestProfile, reqId string, err error) {
	tm := &TokenDBModel{
		ACSignature: toNullString(signature),
	}
	if err = d.db.Where(tm).Find(tm).Error; err != nil {
		return nil, "", err
	}
	if tm.RequestID == "" {
		return nil, "", errors.New("'authorization_code' not found")
	}
	if time.Now().After(tm.ACExpiry) {
		return nil, "", errors.New("'authorization_code' expired")
	}
	return tm.RequestProfile, tm.RequestID, nil
}

func (d *TokenStore) GetProfileWithAccessTokenSign(_ context.Context, signature string) (profile oidcsdk.RequestProfile, reqId string, err error) {
	tm := &TokenDBModel{
		ATSignature: toNullString(signature),
	}
	if err = d.db.Where(tm).Find(tm).Error; err != nil {
		return nil, "", err
	}
	if tm.RequestID == "" {
		return nil, "", errors.New("'access_token' not found")
	}
	if time.Now().After(tm.ATExpiry) {
		return nil, "", errors.New("'access_token' expired")
	}
	return tm.RequestProfile, tm.RequestID, nil
}

func (d *TokenStore) GetProfileWithRefreshTokenSign(_ context.Context, signature string) (profile oidcsdk.RequestProfile, reqId string, err error) {
	tm := &TokenDBModel{
		RTSignature: toNullString(signature),
	}
	if err = d.db.Where(tm).Find(tm).Error; err != nil {
		return nil, "", err
	}
	if tm.RequestID == "" {
		return nil, "", errors.New("'refresh_token' not found")
	}
	if time.Now().After(tm.RTExpiry) {
		return nil, "", errors.New("'refresh_token' expired")
	}
	return tm.RequestProfile, tm.RequestID, nil
}

func (d *TokenStore) InvalidateWithRequestID(_ context.Context, reqID string, what uint8) (err error) {
	tm := &TokenDBModel{
		RequestID: reqID,
	}
	where := d.db.Where(tm)
	err = where.Find(tm).Error
	if err != nil {
		return err
	}
	if what&oidcsdk.ExpireRefreshToken > 0 {
		tm.RTExpiry = time.Now()
	}
	if what&oidcsdk.ExpireAccessToken > 0 {
		tm.ATExpiry = time.Now()
	}
	if what&oidcsdk.ExpireAuthorizationCode > 0 {
		tm.ACExpiry = time.Now()
	}
	return d.db.Save(tm).Error
}

func toNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{
			Valid:  true,
			String: s,
		}
	} else {
		return sql.NullString{Valid: false}
	}
}
