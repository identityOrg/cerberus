package setup

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type AppTemplates struct {
	templates *template.Template
}

func NewAppTemplates() *AppTemplates {
	return &AppTemplates{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func (t *AppTemplates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
