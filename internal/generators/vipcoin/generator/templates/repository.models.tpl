{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

import (
	"time"

	 "{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"
	 "{{.ModuleNameLower}}/internal/api/repository"
)

// tableName database table name.
const tableName = "{{.TableName}}"

type (
	// {{.ListLowerCamel}} type alias for a slice of {{.NameLowerCamel}}.
	{{.ListLowerCamel}} []{{.NameLowerCamel}}

	// {{.NameLowerCamel}} database model.
	{{.NameLowerCamel}} struct {
		{{.DatabaseModel}}
		CreatedAt time.Time   `db:"created_at"`
		UpdatedAt time.Time   `db:"updated_at"`
	}
)

// toDatabase converts domain {{.NameCamel}} to database model.
func toDatabase(domain {{.PackageDomainName}}) {{.NameLowerCamel}} {
	return {{.NameLowerCamel}}{
		{{.FromDomainToDatabase}}
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

// toDomain converts database model to domain model.
func ({{.Reference}} {{.NameLowerCamel}}) toDomain() {{.PackageDomainName}} {
	return {{.PackageDomainName}}{
		{{.FromDatabaseToDomain}}
		CreatedAt: {{.Reference}}.CreatedAt,
		UpdatedAt: {{.Reference}}.UpdatedAt,
	}
}

{{ if .WithPaginationCheck}}

// toDomain converts database model to domain model.
func ({{.Reference}} {{.ListLowerCamel}}) toDomain(total uint64) {{.PackageDomainNameList}} {
	result := make({{.PackageDomainNames}}, 0, len({{.Reference}}))

	for _, {{.NameLowerCamel}} := range {{.Reference}} {
		result = append(result, {{.NameLowerCamel}}.toDomain())
	}

	return {{.PackageDomainNameList}}{
		{{.NamesCamel}}: result,
		Total: total,
	}
}

{{ else }}

// toDomain converts database model to domain model.
func ({{.Reference}} {{.ListLowerCamel}}) toDomain() {{.PackageDomainNames}} {
	result := make({{.PackageDomainNames}}, 0, len({{.Reference}}))

	for _, {{.NameLowerCamel}} := range {{.Reference}} {
		result = append(result, {{.NameLowerCamel}}.toDomain())
	}

	return result
}

{{ end }}