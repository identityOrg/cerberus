package setup

import (
	"github.com/identityOrg/cerberus/impl/session"
	"github.com/identityOrg/cerberus/impl/store"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/compose"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	OAuth2Config   *oidcsdk.Config
	DbConfig       *config.DBConfig
	SecretConfig   *config.SecretConfig
	OrmDB          *gorm.DB
	ClientStore    *store.ClientStore
	UserStore      *store.UserStore
	TokenStore     *store.TokenStore
	KeyStore       *store.KeyStore
	SessionManager *session.Manager
	Manager        oidcsdk.IManager
)

func NewManager(debug bool) (oidcsdk.IManager, error) {
	OAuth2Config = oauth2Config()
	DbConfig = config.NewDBConfig()
	SecretConfig = config.NewSecretConfig()
	var err error
	OrmDB, err = store.NewGOrmDB(DbConfig.Driver, DbConfig.DSN)
	if err != nil {
		return nil, err
	}
	OrmDB.LogMode(debug)
	ClientStore = store.NewClientStore(OrmDB)
	TokenStore = store.NewTokenStore(OrmDB)
	UserStore = store.NewUserStore(OrmDB)
	KeyStore = store.NewKeyStore(OrmDB)
	strategy := strategies.NewDefaultStrategy()
	strategy.HmacKey = SecretConfig.TokenSecret
	sequence := compose.CreateDefaultSequence()
	SessionManager = session.NewManager(OrmDB, SecretConfig.SessionSecret)
	sequence = append(sequence, ClientStore, TokenStore, UserStore, KeyStore, strategy, SessionManager)
	Manager = compose.DefaultManager(OAuth2Config, sequence...)
	compose.SetLoginPageHandler(Manager, RenderLoginPage)
	compose.SetConsentPageHandler(Manager, RenderConsentPage)

	return Manager, nil
}

func oauth2Config() *oidcsdk.Config {
	oauth2Config := oidcsdk.NewConfig("http://localhost:8080")
	oauth2Config.RefreshTokenEntropy = 0
	_ = viper.UnmarshalKey("oauth2", oauth2Config)
	return oauth2Config
}

func ConfigureEcho(e *echo.Echo, manager oidcsdk.IManager) {
	e.GET("/keys", echo.WrapHandler(http.HandlerFunc(manager.ProcessKeysEP)))
	e.GET(oidcsdk.UrlOidcDiscovery, echo.WrapHandler(http.HandlerFunc(manager.ProcessDiscoveryEP)))
	oauth2 := e.Group("/oauth2")
	oauth2.GET("/authorize", echo.WrapHandler(http.HandlerFunc(manager.ProcessAuthorizationEP)))
	oauth2.POST("/token", echo.WrapHandler(http.HandlerFunc(manager.ProcessTokenEP)))
	oauth2.POST("/introspection", echo.WrapHandler(http.HandlerFunc(manager.ProcessIntrospectionEP)))
	oauth2.POST("/revocation", echo.WrapHandler(http.HandlerFunc(manager.ProcessRevocationEP)))
	oauth2.GET("/me", echo.WrapHandler(http.HandlerFunc(manager.ProcessUserInfoEP)))
}
