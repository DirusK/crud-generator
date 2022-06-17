{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
package {{.PackageLower}}

import (
    "github.com/brianvoe/gofakeit/v6"
)

// FakeOption option for fake generator.
type FakeOption func({{.NameLowerCamel}} *{{.NameCamel}})

// Fake{{.NameCamel}} returns new fake {{.NameCamel}}.
func Fake{{.NameCamel}}(opts ...FakeOption) {{.NameCamel}} {
    date := gofakeit.Date()

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
