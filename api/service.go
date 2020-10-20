package api

import "github.com/identityOrg/cerberus/core"

type CerberusAPI struct {
	ScopeClaimStore core.IScopeClaimStoreService
}

func NewCerberusAPI(scopeClaimStore core.IScopeClaimStoreService) *CerberusAPI {
	return &CerberusAPI{ScopeClaimStore: scopeClaimStore}
}
