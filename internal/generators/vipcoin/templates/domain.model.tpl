{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
package {{.PackageLower}}

import (
	"time"
)

// Block which defines all possible fields for {{.NameCamel}}.
const (
	{{.FieldsEnum}}
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
)

type (
	// {{.NamesCamel}} type alias for a slice of {{.NameCamel}}.
	{{.NamesCamel}} []{{.NameCamel}}

	{{ if .WithPaginationCheck}}

	// {{.ListCamel}} describes list of {{.NamesCamel}} with their total number.
	{{.ListCamel}} struct {
		{{ .NamesCamel}} {{ .NamesCamel}}
		Total uint64
	}

	{{end}}

	// {{.NameCamel}} domain model.
	{{.NameCamel}} struct {
		{{.DomainModel}}
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

{{.IsEnumCheck}}