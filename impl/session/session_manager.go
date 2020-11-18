package session

import (
	"github.com/gorilla/sessions"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/identityOrg/oidcsdk"
	"net/http"
)

type Manager struct {
	SessionStore sessions.Store
	SessionName  string
}

func NewManager(secretConfig *config.SecretConfig) *Manager {
	return &Manager{
		SessionStore: sessions.NewCookieStore([]byte(secretConfig.SessionSecret)),
		SessionName:  "cerberus-session",
	}
}

func (m *Manager) RetrieveUserSession(w http.ResponseWriter, r *http.Request) (oidcsdk.ISession, error) {
	sessBack, err := m.SessionStore.Get(r, m.SessionName)
	if err != nil {
		return nil, err
	}
	sess := &DefaultSession{
		s: sessBack,
		r: r,
		w: w,
	}
	return sess, nil
}
