package store

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type KeyDBModel struct {
	gorm.Model
	KID        string `gorm:"column:key_id;size:100;unique_index;not null"`
	Asymmetric bool   `gorm:"column:asymmetric;not null;size:50"`
	Algorithm  string `gorm:"column:algorithm;size:50"`
	Use        string `gorm:"column:use;size:10"`
	PrivateKey string `gorm:"column:private_key;type:lob"`
	PublicKey  string `gorm:"column:public_key;type:lob"`
	SecretKey  string `gorm:"column:secret_key;type:lob"`
}

func (km *KeyDBModel) TableName() string {
	return "keys"
}

func (km *KeyDBModel) SetPrivateKey(key interface{}) error {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		block := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
		memory := pem.EncodeToMemory(&block)
		km.PrivateKey = string(memory)
		km.Asymmetric = true
		return nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return err
		}
		block := pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
		memory := pem.EncodeToMemory(&block)
		km.PrivateKey = string(memory)
		km.Asymmetric = true
		return nil
	}
	return errors.New("unrecognized key type")
}

func (km *KeyDBModel) SetPublicKey(key interface{}) error {
	switch k := key.(type) {
	case *rsa.PublicKey:
		block := pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(k)}
		memory := pem.EncodeToMemory(&block)
		km.PublicKey = string(memory)
		km.Asymmetric = true
		return nil
	case *ecdsa.PublicKey:
		b, err := x509.MarshalPKIXPublicKey(k)
		if err != nil {
			return err
		}
		block := pem.Block{Type: "EC PUBLIC KEY", Bytes: b}
		memory := pem.EncodeToMemory(&block)
		km.PublicKey = string(memory)
		km.Asymmetric = true
		return nil
	}
	return errors.New("unrecognized key type")
}

func (km *KeyDBModel) GetPublicKey() (interface{}, error) {
	if !km.Asymmetric {
		return nil, errors.New(fmt.Sprintf("not an asymmetric key"))
	}
	decoded, _ := pem.Decode([]byte(km.PublicKey))
	if strings.HasPrefix(decoded.Type, "RSA") {
		return x509.ParsePKCS1PublicKey(decoded.Bytes)
	} else if strings.HasPrefix(decoded.Type, "EC") {
		return x509.ParsePKIXPublicKey(decoded.Bytes)
	} else {
		return nil, errors.New(fmt.Sprintf("unrecognized key type %s", decoded.Type))
	}
}

func (km *KeyDBModel) GetPrivateKey() (interface{}, error) {
	if !km.Asymmetric {
		return nil, errors.New(fmt.Sprintf("not an asymmetric key"))
	}
	decoded, _ := pem.Decode([]byte(km.PrivateKey))
	if strings.HasPrefix(decoded.Type, "RSA") {
		return x509.ParsePKCS1PrivateKey(decoded.Bytes)
	} else if strings.HasPrefix(decoded.Type, "EC") {
		return x509.ParseECPrivateKey(decoded.Bytes)
	} else {
		return nil, errors.New(fmt.Sprintf("unrecognized key type %s", decoded.Type))
	}
}
