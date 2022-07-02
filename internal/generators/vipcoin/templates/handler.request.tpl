{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
package {{.PackageLower}}

import "{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"

//go:generate go-validator

// parameter{{.FieldIDCamel}} is required for fetching {{.FieldIDCamel}} parameter from path.
const parameter{{.FieldIDCamel}} = "{{.FieldIDSnake}}"

{{if .WithPagination}}
// defaultLimit describes default quantity of entities to return in response.
const defaultLimit = 20
{{end}}

type (
	// getAllFilter arguments for filtration {{.NamesLowerCamel}} data.
	getAllFilter struct {
		{{.GetAllFilter}}
	}
)
