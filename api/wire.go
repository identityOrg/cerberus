package api

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCerberusAPI,
	wire.Bind(new(ServerInterface), new(*CerberusAPI)),
)
