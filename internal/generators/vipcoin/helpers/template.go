package helpers

import (
	"strings"
	"text/template"

	"crud-generator/pkg/printer"
)

// TagParser parser tag.
const TagParser = "PARSER"

// ExecuteTemplateFromString receives string template and executes it.
func ExecuteTemplateFromString(tmpl string, data interface{}) string {
	parser, err := template.New("parser").Parse(tmpl)
	if err != nil {
		printer.Fatal(TagParser, err, "can't parse template from string")
	}

	var str strings.Builder
	if err = parser.Execute(&str, data); err != nil {
		printer.Fatal(TagParser, err, "can't execute template from string")
	}

	return str.String()
}
