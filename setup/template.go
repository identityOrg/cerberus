package setup

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
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
		box := packr.New("templates", "../templates")
		t := template.New("root")
		for _, s := range box.List() {
			findString, err := box.FindString(s)
			if err != nil {
				fmt.Printf("error in packr box config: %s\n", err.Error())
			} else {
				_, _ = t.New(s).Parse(findString)
			}
		}
		appTemplates = &AppTemplates{
			templates: t,
		}
	}
	return appTemplates
}

func (t *AppTemplates) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
