package store

import (
	"encoding/json"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
	"gopkg.in/square/go-jose.v2"
)

type ClientDBModel struct {
	gorm.Model
	ClientID      string                 `gorm:"column:client_id;unique_index;not null;size:256"`
	ClientSecret  string                 `gorm:"column:client_secret;size:256"`
	Public        bool                   `gorm:"column:public;unique_index;not null"`
	Attributes    map[string]interface{} `gorm:"-"`
	AttributeData string                 `gorm:"column:attributes;type:lob"`
}

func NewClientDBModel() *ClientDBModel {
	return &ClientDBModel{
		Attributes: make(map[string]interface{}),
	}
}

func (c *ClientDBModel) BeforeCreate() error {
	marshal, err := json.Marshal(c.Attributes)
	if err != nil {
		return err
	}
	c.AttributeData = string(marshal)
	return nil
}

func (c *ClientDBModel) AfterFind() (err error) {
	if c.AttributeData != "" {
		var rp map[string]interface{}
		err = json.Unmarshal([]byte(c.AttributeData), &rp)
		if err != nil {
			return err
		}
		c.Attributes = rp
	}
	return
}

func (c *ClientDBModel) TableName() string {
	return "clients"
}

func (c *ClientDBModel) GetID() string {
	return c.ClientID
}

func (c *ClientDBModel) GetSecret() string {
	return c.ClientSecret
}

func (c *ClientDBModel) IsPublic() bool {
	return c.Public
}

func (c *ClientDBModel) GetIDTokenSigningAlg() jose.SignatureAlgorithm {
	signAlgo := c.Attributes["id_token_signed_response_alg"]
	return jose.SignatureAlgorithm(signAlgo.(string))
}

func (c *ClientDBModel) SetIDTokenSigningAlg(alg jose.SignatureAlgorithm) {
	c.Attributes["id_token_signed_response_alg"] = alg
}

func (c *ClientDBModel) GetRedirectURIs() []string {
	redirectUris := c.Attributes["redirect_uris"]
	return c.tryConvertArguments(redirectUris)
}

func (c *ClientDBModel) SetRedirectURIs(uris []string) {
	c.Attributes["redirect_uris"] = uris
}

func (c *ClientDBModel) GetApprovedScopes() oidcsdk.Arguments {
	approvedScopes := c.Attributes["scopes"]
	return c.tryConvertArguments(approvedScopes)
}

func (c *ClientDBModel) SetApprovedScopes(scp oidcsdk.Arguments) {
	c.Attributes["scopes"] = scp
}

func (c *ClientDBModel) GetApprovedGrantTypes() oidcsdk.Arguments {
	asInt := c.Attributes["grant_types"]
	return c.tryConvertArguments(asInt)
}

func (c *ClientDBModel) SetApprovedGrantTypes(gty oidcsdk.Arguments) {
	c.Attributes["grant_types"] = gty
}

func (c *ClientDBModel) GetAttribute(key string) interface{} {
	return c.Attributes[key]
}

func (c *ClientDBModel) SetAttribute(key string, value interface{}) {
	c.Attributes[key] = value
}

func (c *ClientDBModel) tryConvertArguments(asInt interface{}) []string {
	switch k := asInt.(type) {
	case []string:
		return k
	case string:
		return []string{k}
	case []interface{}:
		var scopes []string
		for _, i := range k {
			if e, ok := i.(string); ok {
				scopes = append(scopes, e)
			}
		}
		return scopes
	case interface{}:
		if e, ok := k.(string); ok {
			return []string{e}
		}
	}
	return nil
}
