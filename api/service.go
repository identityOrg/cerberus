package api

import "github.com/identityOrg/cerberus/core"

type CerberusAPI struct {
	ScopeClaimStore  core.IScopeClaimStoreService
	SecretStoreStore core.ISecretStoreService
}

func NewCerberusAPI(
	scopeClaimStore core.IScopeClaimStoreService,
	secretStoreStore core.ISecretStoreService,
) *CerberusAPI {
	return &CerberusAPI{
		ScopeClaimStore:  scopeClaimStore,
		SecretStoreStore: secretStoreStore,
	}
}
