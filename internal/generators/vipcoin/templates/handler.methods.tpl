{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
package {{.PackageLower}}

import (
	"strconv"

	"{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"
	"{{.ModuleNameLower}}/internal/api/services"

	"git.ooo.ua/vipcoin/lib/filter"
	"git.ooo.ua/vipcoin/lib/http/query"
	"git.ooo.ua/vipcoin/lib/http/responder"

	"github.com/gofiber/fiber/v2"
)

//go:generate ifacemaker -f {{.GoFileSnakeWithExtension}} -s Handler -p delivery -i {{.Interface}}HTTP -y "{{.Interface}}HTTP - describe an interface for working with {{.NamesLowerCamel}} over HTTP."

var _ delivery.{{.Interface}}HTTP = &Handler{}

// Handler - define http handler struct for handling {{.NameLowerCamel}} requests.
type Handler struct {
	responder.Responder
	{{.NamesServiceLowerCamel}} services.{{.Interface}}
}

// NewHandler - constructor.
func NewHandler({{.NamesServiceLowerCamel}} services.{{.Interface}}) *Handler {
	return &Handler{
		{{.NamesServiceLowerCamel}}: {{.NamesServiceLowerCamel}},
	}
}

// Get - define http handler method which responds with one {{.NameLowerCamel}} by specified id.
func (h Handler) Get(ctx *fiber.Ctx) error {
	id, err := h.GetCustomParameterID(ctx, parameter{{.FieldIDCamel}})
	if err != nil {
		return err
	}

	result, err := h.{{.NamesServiceLowerCamel}}.Get(
		ctx.Context(),
		filter.NewFilter().SetArgument({{.PackageLower}}.Field{{.FieldIDCamel}}, id))
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toResponse(result))
}

// GetAll - define http handler method which responds with all {{.NamesLowerCamel}}.
func (h Handler) GetAll(ctx *fiber.Ctx) error {
	request := query.NewRequest(ctx, query.WithDefaultCondition(filter.ConditionAND))
	if err := request.SetArgumentsFromStruct(&getAllFilter{}); err != nil {
		return errs.BadRequest{Cause: "invalid filter parameters"}
	}

	if err := request.SetCondition(); err != nil {
		return err
	}

	{{if .WithPaginationCheck}}
	if err := request.SetPagination(defaultLimit); err != nil {
		return err
	}
	{{ end }}
	result, err := h.{{.NamesServiceLowerCamel}}.GetAll(ctx.Context(), request.ToFilter())
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toResponseList(result{{if .WithPaginationCheck}},request.Pagination {{ end }}))
}

// Delete - define http handler method which deletes {{.NameLowerCamel}} by specified id.
func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := h.GetCustomParameterID(ctx, parameter{{.FieldIDCamel}})
	if err != nil {
		return err
	}

	if err = h.{{.NamesServiceLowerCamel}}.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return h.RespondEmpty(ctx, fiber.StatusOK)
}
