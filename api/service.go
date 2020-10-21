package api

import (
	"github.com/identityOrg/cerberus/core"
)

type CerberusAPI struct {
	ScopeClaimStore  core.IScopeClaimStoreService
	SecretStoreStore core.ISecretStoreService
	SPStoreService   core.ISPStoreService
}

func NewCerberusAPI(
	scopeClaimStore core.IScopeClaimStoreService,
	secretStoreStore core.ISecretStoreService,
	spStoreService core.ISPStoreService,
) *CerberusAPI {
	return &CerberusAPI{
		ScopeClaimStore:  scopeClaimStore,
		SecretStoreStore: secretStoreStore,
		SPStoreService:   spStoreService,
	}
}
