{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

type (
	// {{.EnumCamel}} type alias for {{.NameCamel}} {{.EnumLowerCamel}}.
	{{.EnumCamel}} string
)

// Block which defines all possible {{.EnumCamel}} types for {{.NameCamel}}.
const (
	{{.EnumConstants}}
)

// {{.EnumLowerCamel}}Map map for validation.
var {{.EnumLowerCamel}}Map = map[{{.EnumCamel}}]struct{}{
	{{.EnumMap}}
}

// String - method for casting {{.EnumCamel}} to string.
func ({{.Reference}} {{.EnumCamel}}) String() string {
	return string({{.Reference}})
}

// Validate - returns true if {{.EnumCamel}} is valid.
func ({{.Reference}} {{.EnumCamel}}) Validate() bool {
	_, ok := {{.EnumLowerCamel}}Map[{{.Reference}}]
	return ok
}
