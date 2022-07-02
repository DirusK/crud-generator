{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
package {{.PackageLower}}

import (
	"{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"

	"git.ooo.ua/vipcoin/lib/errs"
	"github.com/gofiber/fiber/v2"
)

//go:generate go-validator

// parameter{{.FieldIDCamel}} is required for fetching {{.FieldIDCamel}} parameter from path.
const parameter{{.FieldIDCamel}} = "{{.FieldIDSnake}}"

{{if .WithPagination}}
// defaultLimit describes default quantity of entities to return in response.
const defaultLimit = 20
{{end}}

type (
	// {{.NameLowerCamelRequest}} - request model for creating and updating {{.NameCamel}}
	{{.NameLowerCamelRequest}} struct {
		{{.RequestModel}}
	}

	// getAllFilter arguments for filtration {{.NamesLowerCamel}} data.
	getAllFilter struct {
		{{.FilterModel}}
	}
)

// toDomain() converts request model to domain model.
func ({{.Reference}} {{.NameLowerCamelRequest}}) toDomain() {{.PackageDomainName}} {
	return {{.PackageDomainName}} {
		{{.FromRequestToDomain}}
	}
}

// {{.NameLowerCamelRequest}}FromContext - parses data to the request model from fiber context.
func (h Handler) {{.NameLowerCamelRequest}}FromContext(ctx *fiber.Ctx) ({{.NameLowerCamelRequest}}, error) {
	var request {{.NameLowerCamelRequest}}
	if err := ctx.BodyParser(&request); err != nil {
		return {{.NameLowerCamelRequest}}{}, errs.BadRequest{Cause: "invalid body"}
	}

	if errsList := request.Validate(); len(errsList) != 0 {
		return {{.NameLowerCamelRequest}}{}, errs.FieldsValidation{Errors: errsList}
	}

	return request, nil
}
