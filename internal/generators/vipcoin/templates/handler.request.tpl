package {{.Package}}

//go:generate go-validator

// parameterID is required for fetching ID parameter from path.
const parameterID = "id"

{{if .WithPagination}}
// defaultLimit describes default quantity of entities to return in response.
const defaultLimit = 20
{{end}}

type (
	// getAllFilter arguments for filtration {{.Entities}} data.
	getAllFilter struct {
		{{.FilterFields}}
	}
)
