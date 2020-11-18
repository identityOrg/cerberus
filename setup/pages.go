package setup

import (
	"github.com/identityOrg/oidcsdk"
	"net/http"
)

type CerberusPageResponseHandler struct {
	template *AppTemplates
}

func NewCerberusPageResponseHandler(template *AppTemplates) *CerberusPageResponseHandler {
	return &CerberusPageResponseHandler{template: template}
}

func (c *CerberusPageResponseHandler) DisplayLogoutConsentPage(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c *CerberusPageResponseHandler) DisplayLogoutStatusPage(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c *CerberusPageResponseHandler) DisplayErrorPage(err error, writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set(oidcsdk.HeaderContentType, oidcsdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = c.template.Render(writer, "error.html", err, nil)
}

func (c *CerberusPageResponseHandler) DisplayLoginPage(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set(oidcsdk.HeaderContentType, oidcsdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = c.template.Render(writer, "login.html", r.URL.String(), nil)
}

func (c *CerberusPageResponseHandler) DisplayConsentPage(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set(oidcsdk.HeaderContentType, oidcsdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = c.template.Render(writer, "consent.html", r.URL.String(), nil)
}
