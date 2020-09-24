//+build wireinject

package setup

import (
	"context"
	"github.com/google/wire"
	core "github.com/identityOrg/cerberus-core"
	"github.com/identityOrg/cerberus/impl/handlers"
	"github.com/identityOrg/cerberus/impl/session"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl"
	"github.com/labstack/echo/v4"
)

var AppDependency = wire.NewSet(
	NewGormDB,
	NewEchoServer,
	NewAppTemplates,
	core.ProviderSet,
	impl.DefaultManagerSet,
	impl.DefaultProcessorSet,
	config.NewCoreConfig,
	config.NewSDKConfig,
	config.NewDBConfig,
	config.NewSecretConfig,
	config.NewServerConfig,
	handlers.NewLoginHandler,
	session.NewManager,
	wire.Bind(new(oidcsdk.ISessionManager), new(*session.Manager)),
	NewCrypto,
	wire.Bind(new(core.ITextEncrypts), new(*Crypto)),
	wire.Bind(new(core.ITextDecrypts), new(*Crypto)),
)

func CreateEchoServer() (*echo.Echo, error) {
	wire.Build(AppDependency)
	return nil, nil
}

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
