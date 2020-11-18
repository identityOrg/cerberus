package handlers

import (
	"github.com/identityOrg/cerberus/impl/session"
	"github.com/identityOrg/oidcsdk"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LoginHandler struct {
	UserStore      oidcsdk.IUserStore
	SessionManager *session.Manager
}

func NewLoginHandler(userStore oidcsdk.IUserStore, sessionManager *session.Manager) *LoginHandler {
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
		userSession, err := l.SessionManager.RetrieveUserSession(context.Response(), context.Request())
		if err != nil {
			return err
		} else {
			sess := userSession.(*session.DefaultSession)
			sess.SetAttribute(session.UsernameAttribute, username)
			now := time.Now()
			sess.SetAttribute(session.LoginTimeAttribute, &now)

			err := sess.Save()
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
