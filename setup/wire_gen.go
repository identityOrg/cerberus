// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package setup

import (
	"context"
	"github.com/google/wire"
	"github.com/identityOrg/cerberus-core"
	"github.com/identityOrg/cerberus/impl/handlers"
	"github.com/identityOrg/cerberus/impl/session"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl"
	"github.com/identityOrg/oidcsdk/impl/factories"
	"github.com/identityOrg/oidcsdk/impl/manager"
	"github.com/identityOrg/oidcsdk/impl/processors"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"github.com/identityOrg/oidcsdk/impl/writers"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/identityOrg/cerberus/packrd"
)

// Injectors from wire.go:

func CreateEchoServer() (*echo.Echo, error) {
	serverConfig := config.NewServerConfig()
	setupAppTemplates := NewAppTemplates()
	oidcsdkConfig := config.NewSDKConfig()
	defaultRequestContextFactory := factories.NewDefaultRequestContextFactory()
	defaultErrorWriter := writers.NewDefaultErrorWriter()
	defaultResponseWriter := writers.NewDefaultResponseWriter()
	dbConfig := config.NewDBConfig()
	db, err := NewGormDB(dbConfig)
	if err != nil {
		return nil, err
	}
	secretConfig := config.NewSecretConfig()
	sessionManager := session.NewManager(db, secretConfig)
	secretStoreServiceImpl := core.NewSecretStoreServiceImpl(db)
	tokenStoreServiceImpl := core.NewTokenStoreServiceImpl(db)
	coreConfig := config.NewCoreConfig()
	userStoreServiceImpl := core.NewUserStoreServiceImpl(db, coreConfig)
	defaultStrategy := strategies.NewDefaultStrategy(secretStoreServiceImpl, oidcsdkConfig)
	defaultBearerUserAuthProcessor := processors.NewDefaultBearerUserAuthProcessor(tokenStoreServiceImpl, userStoreServiceImpl, defaultStrategy)
	crypto := NewCrypto()
	spStoreServiceImpl := core.NewSPStoreServiceImpl(db, crypto, crypto)
	defaultClientAuthenticationProcessor := processors.NewDefaultClientAuthenticationProcessor(spStoreServiceImpl)
	defaultGrantTypeValidator := processors.NewDefaultGrantTypeValidator()
	defaultResponseTypeValidator := processors.NewDefaultResponseTypeValidator()
	defaultAccessCodeValidator := processors.NewDefaultAccessCodeValidator(tokenStoreServiceImpl, defaultStrategy)
	defaultRefreshTokenValidator := processors.NewDefaultRefreshTokenValidator(defaultStrategy, tokenStoreServiceImpl)
	defaultStateValidator := processors.NewDefaultStateValidator(oidcsdkConfig)
	defaultPKCEValidator := processors.NewDefaultPKCEValidator(oidcsdkConfig)
	defaultRedirectURIValidator := processors.NewDefaultRedirectURIValidator()
	defaultAudienceValidationProcessor := processors.NewDefaultAudienceValidationProcessor()
	defaultScopeValidator := processors.NewDefaultScopeValidator()
	defaultUserValidator := processors.NewDefaultUserValidator(userStoreServiceImpl, spStoreServiceImpl, oidcsdkConfig)
	defaultClaimProcessor := processors.NewDefaultClaimProcessor(userStoreServiceImpl)
	defaultTokenIntrospectionProcessor := processors.NewDefaultTokenIntrospectionProcessor(tokenStoreServiceImpl, defaultStrategy, defaultStrategy)
	defaultTokenRevocationProcessor := processors.NewDefaultTokenRevocationProcessor(tokenStoreServiceImpl, defaultStrategy, defaultStrategy)
	defaultAuthCodeIssuer := processors.NewDefaultAuthCodeIssuer(defaultStrategy, oidcsdkConfig)
	defaultAccessTokenIssuer := processors.NewDefaultAccessTokenIssuer(defaultStrategy, oidcsdkConfig)
	defaultIDTokenIssuer := processors.NewDefaultIDTokenIssuer(defaultStrategy, oidcsdkConfig)
	defaultRefreshTokenIssuer := processors.NewDefaultRefreshTokenIssuer(defaultStrategy, oidcsdkConfig)
	defaultTokenPersister := processors.NewDefaultTokenPersister(tokenStoreServiceImpl, userStoreServiceImpl, oidcsdkConfig)
	v := processors.NewProcessorSequence(defaultBearerUserAuthProcessor, defaultClientAuthenticationProcessor, defaultGrantTypeValidator, defaultResponseTypeValidator, defaultAccessCodeValidator, defaultRefreshTokenValidator, defaultStateValidator, defaultPKCEValidator, defaultRedirectURIValidator, defaultAudienceValidationProcessor, defaultScopeValidator, defaultUserValidator, defaultClaimProcessor, defaultTokenIntrospectionProcessor, defaultTokenRevocationProcessor, defaultAuthCodeIssuer, defaultAccessTokenIssuer, defaultIDTokenIssuer, defaultRefreshTokenIssuer, defaultTokenPersister)
	options := &manager.Options{
		RequestContextFactory: defaultRequestContextFactory,
		ErrorWriter:           defaultErrorWriter,
		ResponseWriter:        defaultResponseWriter,
		UserSessionManager:    sessionManager,
		SecretStore:           secretStoreServiceImpl,
		Sequence:              v,
	}
	defaultManager := manager.NewDefaultManager(oidcsdkConfig, options)
	loginHandler := handlers.NewLoginHandler(userStoreServiceImpl, sessionManager)
	echoEcho := NewEchoServer(serverConfig, setupAppTemplates, defaultManager, loginHandler)
	return echoEcho, nil
}

// wire.go:

var AppDependency = wire.NewSet(
	NewGormDB,
	NewEchoServer,
	NewAppTemplates, core.ProviderSet, impl.DefaultManagerSet, impl.DefaultProcessorSet, config.NewCoreConfig, config.NewSDKConfig, config.NewDBConfig, config.NewSecretConfig, config.NewServerConfig, handlers.NewLoginHandler, session.NewManager, wire.Bind(new(oidcsdk.ISessionManager), new(*session.Manager)), NewCrypto, wire.Bind(new(core.ITextEncrypts), new(*Crypto)), wire.Bind(new(core.ITextDecrypts), new(*Crypto)),
)

type Crypto struct{}

func NewCrypto() *Crypto {
	return &Crypto{}
}

func (*Crypto) DecryptText(ctx context.Context, cypherText string) (text string, err error) {
	return cypherText, nil
}

func (*Crypto) EncryptText(ctx context.Context, text string) (cypherText string, err error) {
	return text, nil
}