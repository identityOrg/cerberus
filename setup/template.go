package setup

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"os"
)

type AppTemplates struct {
	templates *template.Template
}

var appTemplates *AppTemplates

func NewAppTemplates() *AppTemplates {
	if appTemplates == nil {
		templateBox := rice.MustFindBox("../templates")
		t := template.New("root")
		_ = templateBox.Walk("/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				tmplString := templateBox.MustString(info.Name())
				_, _ = t.New(info.Name()).Parse(tmplString)
			}
			return nil
		})
		appTemplates = &AppTemplates{
			templates: t,
		}
	}
	return appTemplates
}

func (t *AppTemplates) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
