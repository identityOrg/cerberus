package store

import (
	"context"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
)

type ClientStore struct {
	db *gorm.DB
}

func NewClientStore(db *gorm.DB) *ClientStore {
	return &ClientStore{db: db}
}

func (c *ClientStore) AutoMigrate() {
	c.db.AutoMigrate(ClientDBModel{})
}

func (c *ClientStore) GetClient(_ context.Context, clientID string) (oidcsdk.IClient, error) {
	client := &ClientDBModel{
		ClientID: clientID,
	}
	err := c.db.Where(client).Find(client).Error
	if err != nil {
		return nil, err
	} else {
		return client, nil
	}
}

func (c *ClientStore) FetchClientProfile(context.Context, string) oidcsdk.RequestProfile {
	panic("implement me")
}
