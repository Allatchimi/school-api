package helpers

import (
	"os"
	"text/template"
)

// RenderTemplate renders a template file with provided data.
func RenderTemplate(filePath, templateContent string, data any) error {
	tmpl, err := template.New("tpl").Parse(templateContent)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
