// Package templates defines project temps
package templates

import (
	"bytes"
	"fmt"
	"html/template"
)

type TemplateManager struct {
	templates map[string]*template.Template
}

func New() *TemplateManager {
	templates := make(map[string]*template.Template)
	t := template.New("confirmation_letter")
	temp, err := t.Parse(ConfirmationLetterFS)
	if err != nil {
		panic(err)
	}

	templates["confirmation_letter"] = temp

	return &TemplateManager{
		templates: templates,
	}
}

func (tm *TemplateManager) Render(name string, fields map[string]string) (string, error) {
	tmpl, exists := tm.templates[name]
	if !exists {
		return "", fmt.Errorf("template %s not found", name)
	}

	var result bytes.Buffer
	err := tmpl.Execute(&result, fields)
	if err != nil {
		return "", fmt.Errorf("failed to execute on template: %w", err)
	}

	return result.String(), nil
}
