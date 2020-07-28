package session

import (
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/util"
	"strings"
	"time"
)

type Session struct {
	Username  string
	Scope     string
	LoginTime *time.Time
}

func (d Session) GetScope() string {
	return d.Scope
}

func (d Session) IsLoginDone() bool {
	return true
}

func (d Session) IsConsentSubmitted() bool {
	return true
}

func (d Session) GetApprovedScopes() oidcsdk.Arguments {
	return util.RemoveEmpty(strings.Split(d.Scope, " "))
}

func (d Session) GetUsername() string {
	return d.Username
}

func (d Session) GetLoginTime() *time.Time {
	return d.LoginTime
}
