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
	return redirectUris.([]string)
}

func (c *ClientDBModel) SetRedirectURIs(uris []string) {
	c.Attributes["redirect_uris"] = uris
}

func (c *ClientDBModel) GetApprovedScopes() oidcsdk.Arguments {
	approvedScopes := c.Attributes["scopes"]
	return approvedScopes.([]string)
}

func (c *ClientDBModel) SetApprovedScopes(scp oidcsdk.Arguments) {
	c.Attributes["scopes"] = scp
}

func (c *ClientDBModel) GetApprovedGrantTypes() oidcsdk.Arguments {
	approvedScopes := c.Attributes["grant_types"]
	return approvedScopes.([]string)
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
