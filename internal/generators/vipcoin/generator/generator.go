package generator

import (
	"text/template"

	"crud-generator-gui/internal/generators"
	vipcoin "crud-generator-gui/internal/generators/vipcoin/models"
	"crud-generator-gui/internal/models"
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
		Template: template.New("generator"),
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
