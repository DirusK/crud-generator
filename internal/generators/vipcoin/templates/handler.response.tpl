package {{.Package}}

type (
	// {{.Entity}} - response model for {{.Entity}}.
	{{.Entity}} struct {
		{{.ResponseModel}}
	}

{{if .WithPagination}}

	// {{.EntityListPrivate}} - response model for list of {{.EntitiesPrivate}}.
		{{.EntityListPrivate}} struct {
		{{.EntitiesPublic}} []{{.Entity}} `json:"{{.CamelNames}}"`
	}

	// {{.EntityListPrivate}}Response - response model with pagination and data.
	{{.EntityListPrivate}}Response struct {
		Result {{.EntityListPrivate}} `json:"result"`
		Pagination query.PaginationRequest `json:"pagination"`
	}

{{ else }}

	// {{.EntityListPrivate}} - response model for list of {{.EntitiesPrivate}}.
	{{.EntityListPrivate}} []{{.Entity}}

{{ end }}
)

// toResponse converts domain {{.DomainPkg}} to response {{.Entity}} model.
func toResponse(domain {{.DomainPkg}}) {{.Entity}} {
	return {{.Entity}}{
		{{.FromDomain}}
	}
}
{{if .WithPagination}}
// toResponseList converts domain {{.DomainPkgList}} to response {{.EntityListPrivate}} model.
func toResponseList(list {{.DomainPkgList}}, pagination query.PaginationRequest) {{.EntityListPrivate}}Response {
	response := make([]{{.Entity}}, 0, len(list.{{.EntitiesPublic}}))
	for _, ent := range list.{{.EntitiesPublic}} {
		response = append(response, toResponse(ent))
	}

	return {{.EntityListPrivate}}Response {
		Result: {{.EntityListPrivate}}{
			{{.EntitiesPublic}}: response,
		},
		Pagination: query.PaginationRequest{
			Offset: pagination.Offset,
			Limit:  pagination.Limit,
			Total:  list.Total,
		},
	}
}
{{ else }}

// toResponseList converts domain {{.DomainPkgList}} to response {{.EntityListPrivate}} model.
func toResponseList(domain {{.DomainPkgList}}) {{.EntityListPrivate}} {
	response := make({{.EntityListPrivate}}, 0, len(domain))
	for _, vote := range domain {
		response = append(response, toResponse(vote))
	}

	return response
}

{{ end }}