package setup

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type AppTemplates struct {
	templates *template.Template
}

var appTemplates *AppTemplates

func NewAppTemplates() *AppTemplates {
	if appTemplates == nil {
		appTemplates = &AppTemplates{
			templates: template.Must(template.ParseGlob("templates/*.html")),
		}
	}
	return appTemplates
}

func (t *AppTemplates) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
