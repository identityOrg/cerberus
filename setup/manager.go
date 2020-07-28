package setup

import (
	"github.com/identityOrg/cerberus/impl/session"
	"github.com/identityOrg/cerberus/impl/store"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/compose"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func NewManager() (oidcsdk.IManager, error) {
	config := oidcsdk.NewConfig("http://localhost:8080")
	config.RefreshTokenEntropy = 0
	_ = viper.UnmarshalKey("oauth2", config)
	db, err := store.NewGOrmDB("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	clientStore := store.NewClientStore(db)
	tokenStore := store.NewTokenStore(db)
	userStore := store.NewUserStore(db)
	keyStore := store.NewKeyStore(db)
	strategy := strategies.NewDefaultStrategy()
	sequence := compose.CreateDefaultSequence()
	sessionManager := session.NewManager(db, "session-secret-key")
	sequence = append(sequence, clientStore, tokenStore, userStore, keyStore, strategy, sessionManager)
	manager := compose.DefaultManager(config, sequence...)
	compose.SetLoginPageHandler(manager, RenderLoginPage)
	compose.SetConsentPageHandler(manager, RenderConsentPage)

	return manager, nil
}

func ConfigureEcho(e *echo.Echo, manager oidcsdk.IManager) {
	e.GET("/keys", echo.WrapHandler(http.HandlerFunc(manager.ProcessKeysEP)))
	e.GET(oidcsdk.UrlOidcDiscovery, echo.WrapHandler(http.HandlerFunc(manager.ProcessDiscoveryEP)))
	oauth2 := e.Group("/oauth2")
	oauth2.GET("/authorize", echo.WrapHandler(http.HandlerFunc(manager.ProcessAuthorizationEP)))
	oauth2.POST("/token", echo.WrapHandler(http.HandlerFunc(manager.ProcessTokenEP)))
	oauth2.POST("/introspection", echo.WrapHandler(http.HandlerFunc(manager.ProcessIntrospectionEP)))
	oauth2.POST("/revocation", echo.WrapHandler(http.HandlerFunc(manager.ProcessRevocationEP)))
}
