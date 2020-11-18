package session

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/util"
	"net/http"
	"strings"
	"time"
)

func init() {
	gob.Register(&time.Time{})
}

const (
	ScopeAttribute        = "scope"
	LoginTimeAttribute    = "login-time"
	UsernameAttribute     = "username"
	LoginFlashAttribute   = "login"
	ConsentFlashAttribute = "consent"
)

type DefaultSession struct {
	s *sessions.Session
	r *http.Request
	w http.ResponseWriter
}

func (d DefaultSession) SetAttribute(name string, value interface{}) {
	d.s.Values[name] = value
}

func (d DefaultSession) GetAttribute(name string) interface{} {
	return d.s.Values[name]
}

func (d DefaultSession) Save() error {
	return d.s.Save(d.r, d.w)
}

func (d DefaultSession) Logout() {
	d.s.Options.MaxAge = -1
}

func (d DefaultSession) GetScope() string {
	scope := d.s.Values[ScopeAttribute]
	if scope != nil {
		return scope.(string)
	}
	return ""
}

func (d DefaultSession) IsLoginDone() bool {
	logins := d.s.Flashes(LoginFlashAttribute)
	if len(logins) > 0 {
		return true
	}
	return false
}

func (d DefaultSession) IsConsentSubmitted() bool {
	consents := d.s.Flashes(ConsentFlashAttribute)
	if len(consents) > 0 {
		return true
	}
	return false
}

func (d DefaultSession) GetApprovedScopes() oidcsdk.Arguments {
	return util.RemoveEmpty(strings.Split(d.GetScope(), " "))
}

func (d DefaultSession) GetUsername() string {
	userName := d.s.Values[UsernameAttribute]
	if userName != nil {
		return userName.(string)
	}
	return ""
}

func (d DefaultSession) GetLoginTime() *time.Time {
	loginTime := d.s.Values[LoginTimeAttribute]
	if loginTime != nil {
		return loginTime.(*time.Time)
	}
	return nil
}
