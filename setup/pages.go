package setup

import (
	"github.com/identityOrg/oidcsdk"
	"net/http"
)

func RenderLoginPage(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set(oidcsdk.HeaderContentType, oidcsdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = NewAppTemplates().Render(writer, "login.html", r, nil)
}

func RenderConsentPage(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set(oidcsdk.HeaderContentType, oidcsdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = NewAppTemplates().Render(writer, "consent.html", r, nil)
}
