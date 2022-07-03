{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

import (
    "time"

    "github.com/brianvoe/gofakeit/v6"
)

// FakeOption option for fake generator.
type FakeOption func({{.NameLowerCamel}} *{{.NameCamel}})

// Fake{{.NameCamel}} returns new fake {{.NameCamel}}.
func Fake{{.NameCamel}}(opts ...FakeOption) {{.NameCamel}} {
    date := time.Now()

    {{.NameLowerCamel}} := {{.NameCamel}}{
        {{.FakeModel}}
        CreatedAt:   date,
        UpdatedAt:   date,
    }

    for _, opt := range opts {
        opt(&{{.NameLowerCamel}})
    }

    return {{.NameLowerCamel}}
}

{{.FakeOptions}}
