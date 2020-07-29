package handlers

import (
	"github.com/identityOrg/cerberus/impl/session"
	"github.com/identityOrg/cerberus/impl/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LoginHandler struct {
	UserStore      *store.UserStore
	SessionManager *session.Manager
}

func NewLoginHandler(userStore *store.UserStore, sessionManager *session.Manager) *LoginHandler {
	return &LoginHandler{UserStore: userStore, SessionManager: sessionManager}
}

func (l *LoginHandler) Use(e *echo.Echo) {
	e.POST("/login", l.AuthenticateUser)
}

func (l *LoginHandler) AuthenticateUser(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")
	request := context.FormValue("request")

	err := l.UserStore.Authenticate(nil, username, []byte(password))
	if err != nil {
		return context.Render(http.StatusOK, "login.html", request)
	} else {
		userSession, err := l.SessionManager.RetrieveUserSession(context.Request())
		if err != nil {
			return err
		} else {
			s := userSession.(*session.Session)
			s.Username = username
			now := time.Now()
			s.LoginTime = &now

			err := l.SessionManager.StoreUserSession(context.Response(), context.Request(), userSession)
			if err != nil {
				return err
			}
			if request != "" {
				return context.Redirect(http.StatusFound, request)
			} else {
				return context.Redirect(http.StatusFound, "/")
			}
		}
	}
}
