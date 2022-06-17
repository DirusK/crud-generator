package models

import (
	"strings"
	"text/template"

	"crud-generator-gui/pkg/printer"
)

func executeTemplateFromString(tmpl string, data any) string {
	parser, err := template.New("parser").Parse(tmpl)
	if err != nil {
		printer.Fatal("PARSER", err, "can't parse template from string")
	}

	var str strings.Builder
	if err = parser.Execute(&str, data); err != nil {
		printer.Fatal("PARSER", err, "can't execute template from string")
	}

	return str.String()
}
