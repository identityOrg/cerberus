package store

import (
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"gopkg.in/square/go-jose.v2"
)

type KeyStore struct {
	db *gorm.DB
}

func NewKeyStore(db *gorm.DB) *KeyStore {
	return &KeyStore{db: db}
}

func (k *KeyStore) AutoMigrate() {
	k.db.AutoMigrate(KeyDBModel{})
}

func (k *KeyStore) GetAllSecrets() *jose.JSONWebKeySet {
	keySet := &jose.JSONWebKeySet{}
	var keys []KeyDBModel
	k.db.Find(&keys)
	for _, key := range keys {
		jwk := jose.JSONWebKey{
			KeyID:     key.KID,
			Algorithm: key.Algorithm,
			Use:       key.Use,
		}
		var err error
		if key.Asymmetric && key.PrivateKey != "" {
			jwk.Key, err = key.GetPrivateKey()
		} else {
			jwk.Key, err = key.GetPublicKey()
		}
		if err != nil {
			log.Errorf("failed to read key from DB with id %s", key.KID)
		} else {
			keySet.Keys = append(keySet.Keys, jwk)
		}
	}
	return keySet
}
