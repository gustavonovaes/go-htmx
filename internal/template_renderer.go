package internal

import (
	"embed"
	"io"
	"text/template"
)

type TemplateRenderer struct {
	Templates *template.Template
}

var (
	//go:embed */views/*.html
	embedFS embed.FS
)

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func (t *TemplateRenderer) GetTemplates() []string {
	var templates []string
	for _, tmpl := range t.Templates.Templates() {
		templates = append(templates, tmpl.Name())
	}
	return templates
}

func NewTemplateRenderer(globPattern ...string) (*TemplateRenderer, error) {
	template, err := template.ParseFS(
		embedFS,
		globPattern...,
	)

	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{
		template,
	}, nil
}
