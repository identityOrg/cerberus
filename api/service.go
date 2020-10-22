package api

import (
	"github.com/identityOrg/cerberus/core"
)

type CerberusAPI struct {
	ScopeClaimStore  core.IScopeClaimStoreService
	SecretStoreStore core.ISecretStoreService
	SPStoreService   core.ISPStoreService
	UserStoreService core.IUserStoreService
}

func NewCerberusAPI(
	scopeClaimStore core.IScopeClaimStoreService,
	secretStoreStore core.ISecretStoreService,
	spStoreService core.ISPStoreService,
	userStoreService core.IUserStoreService,
) *CerberusAPI {
	return &CerberusAPI{
		ScopeClaimStore:  scopeClaimStore,
		SecretStoreStore: secretStoreStore,
		SPStoreService:   spStoreService,
		UserStoreService: userStoreService,
	}
}
