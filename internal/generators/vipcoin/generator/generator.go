package generator

import (
	"embed"
	"text/template"

	"crud-generator/internal/generators"
	vipcoin "crud-generator/internal/generators/vipcoin/models"
	"crud-generator/internal/models"
	"crud-generator/pkg/printer"
)

//go:embed templates/*.tpl
var templates embed.FS

const (
	tagVipcoinGenerator = "VIPCOIN GENERATOR"
	templatesPath       = "templates"
)

// Generator main structure.
type Generator struct {
	Settings models.Settings
	Entity   vipcoin.Entity
	Template *template.Template
}

// NewGenerator inits VipCoin generator.
func NewGenerator(entity models.Entity, settings models.Settings) generators.Generator {
	return Generator{
		Template: loadTemplates(templates),
		Settings: settings,
		Entity: vipcoin.NewEntity(
			settings.ModuleName,
			entity.Copyright,
			entity.Name,
			entity.Package,
			entity.Table,
			entity.WithPagination,
			entity.Fields,
		),
	}
}

// loadTemplates reads template directory.
func loadTemplates(templatesFS embed.FS) *template.Template {
	dir, err := templatesFS.ReadDir(templatesPath)
	if err != nil {
		printer.Fatal(tagVipcoinGenerator, err, "read fs with templates")
	}

	tmpl := &template.Template{}

	for _, file := range dir {
		// supports only first level dir templates
		if file.IsDir() {
			continue
		}

		data, err := templatesFS.ReadFile(templatesPath + "/" + file.Name())
		if err != nil {
			printer.Fatal(tagVipcoinGenerator, err, "read template file")
		}

		tmpl, err = tmpl.New(file.Name()).Parse(string(data))
		if err != nil {
			printer.Fatal(tagVipcoinGenerator, err, "parse template file")
		}
	}

	return tmpl
}
