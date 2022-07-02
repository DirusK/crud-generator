{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
package {{.PackageLower}}

import (
	"git.ooo.ua/vipcoin/lib/http/query"

	"{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"
)

type (
	// {{.NameLowerCamel}} - response model for {{.NameCamel}}.
	{{.NameLowerCamel}} struct {
		{{.ResponseModel}}
	}

{{ if .WithPagination }}

	// {{.ListLowerCamel}} - response model for list of {{.NamesCamel}}.
		{{.ListLowerCamel}} struct {
		{{.NamesCamel}} []{{.NameLowerCamel}} `json:"{{.NamesSnake}}"`
	}

	// {{.NamesLowerCamelResponse}} - response model with pagination and data.
	{{.NamesLowerCamelResponse}} struct {
		Data {{.ListLowerCamel}} `json:"data"`
		Pagination query.PaginationRequest `json:"pagination"`
	}

{{ else }}

	// {{.ListLowerCamel}} - response model for list of {{.NamesCamel}}.
	{{.ListLowerCamel}} []{{.NameLowerCamel}}

{{ end }}
)

// toResponse converts domain {{.PackageDomainName}} to response {{.NameLowerCamel}} model.
func toResponse(domain {{.PackageDomainName}}) {{.NameLowerCamel}} {
	return {{.NameLowerCamel}}{
		{{.FromDomainToResponse}}
	}
}
{{ if .WithPagination }}
// toResponseList converts domain {{.ListCamel}} to response {{.NamesLowerCamelResponse}} model.
func toResponseList(list {{.PackageDomainNameList}}, pagination query.PaginationRequest) {{.NamesLowerCamelResponse}} {
	response := make([]{{.NameLowerCamel}}, 0, len(list.{{.NamesCamel}}))
	for _, {{.NameLowerCamel}} := range list.{{.NamesCamel}} {
		response = append(response, toResponse({{.NameLowerCamel}}))
	}

	return {{.NamesLowerCamelResponse}} {
		Data: {{.ListLowerCamel}}{
			{{.NamesCamel}}: response,
		},
		Pagination: query.PaginationRequest{
			Offset: pagination.Offset,
			Limit:  pagination.Limit,
			Total:  list.Total,
		},
	}
}

{{ else }}

// toResponseList converts domain {{.ListCamel}} to response {{.ListLowerCamel}} model.
func toResponseList(domain {{.PackageDomainNames}}) {{.ListLowerCamel}} {
	response := make({{.ListLowerCamel}}, 0, len(domain))
	for _, {{.NameLowerCamel}} := range domain {
		response = append(response, toResponse({{.NameLowerCamel}}))
	}

	return response
}

{{ end }}