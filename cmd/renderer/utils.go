package renderer

import (
	"html/template"
)

func NewTemplateRenderer() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
