package {{.Package}}

import (
	"strconv"

	"{{.Project}}/internal/api/domain/{{.Package}}"
	"{{.Project}}/internal/api/services"

	"git.ooo.ua/vipcoin/lib/filter"
	"git.ooo.ua/vipcoin/lib/http/query"
	"git.ooo.ua/vipcoin/lib/http/responder"

	"github.com/gofiber/fiber/v2"
)

//go:generate ifacemaker -f {{.CurrentFile}} -s Handler -p delivery -i {{.Interface}}HTTP -y "{{.Interface}}HTTP - describe an interface for working with {{.Entities}} over HTTP."

var _ delivery.{{.Interface}}HTTP = &Handler{}

// Handler - define http handler struct for handling {{.Entity}} requests.
type Handler struct {
	responder.Responder
	{{.ServiceName}} services.{{.Interface}}
}

// NewHandler - constructor.
func NewHandler({{.Entities}} services.{{.Interface}}) *Handler {
	return &Handler{
		{{.ServiceName}}: {{.Entities}},
	}
}

// Get - define http handler method which responds with one {{.Entity}} by specified id.
func (h Handler) Get(ctx *fiber.Ctx) error {
	id, err := h.GetCustomParameterID(ctx, parameterID)
	if err != nil {
		return err
	}


	result, err := h.{{.ServiceName}}.Get(
		ctx.Context(),
		filter.NewFilter().SetArgument({{.Package}}.FieldID, id))
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toResponse(result))
}

// GetAll - define http handler method with all {{.Entities}}.
func (h Handler) GetAll(ctx *fiber.Ctx) error {
	request := query.NewRequest(ctx, query.WithDefaultCondition(filter.ConditionAND))
	if err := request.SetArgumentsFromStruct(&getAllFilter{}); err != nil {
		return errs.BadRequest{Cause: "invalid filter parameters"}
	}

	if err := request.SetCondition(); err != nil {
		return err
	}

	{{if .WithPagination}}
	if err := request.SetPagination(defaultLimit); err != nil {
		return err
	}
	{{ end }}
	result, err := h.{{.ServiceName}}.GetAll(ctx.Context(), request.ToFilter())
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toResponseList(result{{if .WithPagination}},request.Pagination {{ end }}))
}

// Delete - define http handler method which deletes {{.Entity}} by specified id.
func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := h.GetCustomParameterID(ctx, parameterID)
	if err != nil {
		return err
	}

	if err = h.{{.ServiceName}}.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return h.RespondEmpty(ctx, fiber.StatusOK)
}
