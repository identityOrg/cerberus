package session

import (
	"github.com/gorilla/sessions"
	"github.com/identityOrg/oidcsdk"
	"github.com/jinzhu/gorm"
	"github.com/wader/gormstore"
	"net/http"
	"time"
)

type Manager struct {
	SessionStore sessions.Store
	SessionName  string
}

func NewManager(db *gorm.DB, secretKey string) *Manager {
	return &Manager{
		SessionStore: gormstore.New(db, []byte(secretKey)),
		SessionName:  "cerberus-session",
	}
}

func (m *Manager) RetrieveUserSession(r *http.Request) (oidcsdk.ISession, error) {
	sessBack, err := m.SessionStore.Get(r, m.SessionName)
	if err != nil {
		return nil, err
	}
	sess := &Session{}
	userName := sessBack.Values["username"]
	scope := sessBack.Values["scope"]
	loginTime := sessBack.Values["login-time"]
	if userName != nil {
		sess.Username = userName.(string)
	}
	if loginTime != nil {
		i := loginTime.(int64)
		unix := time.Unix(i, 0)
		sess.LoginTime = &unix
	}
	if scope != nil {
		sess.Scope = scope.(string)
	}
	return sess, nil
}

func (m *Manager) StoreUserSession(w http.ResponseWriter, r *http.Request, sess oidcsdk.ISession) error {
	sessBack, err := m.SessionStore.Get(r, m.SessionName)
	if err != nil {
		return err
	}
	sessBack.Values["username"] = sess.GetUsername()
	sessBack.Values["scope"] = sess.GetScope()
	sessBack.Values["login-time"] = sess.GetLoginTime().Unix()

	return m.SessionStore.Save(r, w, sessBack)
}
