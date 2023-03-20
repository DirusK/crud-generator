{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

import (
	"context"
	"errors"

	"git.ooo.ua/vipcoin/lib/errs"
	"git.ooo.ua/vipcoin/lib/filter"
	"git.ooo.ua/vipcoin/lib/log"

	"{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"
	"{{.ModuleNameLower}}/internal/api/repository"
	"{{.ModuleNameLower}}/internal/api/services"
)

//go:generate ifacemaker -f {{.GoFileSnakeWithExtension}} -s Service -p services -i {{.Interface}} -y "{{.Interface}} - describe an interface for working with {{.NamesLowerCamel}} business logic."

var _ services.{{.Interface}} = &Service{}

// Service - defines {{.NameLowerCamel}} service struct.
type Service struct {
    datastore *repository.Datastore
	logger    log.Logger
}

// NewService - constructor.
func NewService(datastore *repository.Datastore, logger log.Logger) *Service {
	return &Service{
		datastore: datastore,
		logger:    logger.With("{{.NamesLowerSpace}} service"),
	}
}

// Create - method for creating {{.NameLowerSpace}}.
func (s Service) Create(ctx context.Context, {{.NameLowerCamel}} {{.PackageDomainName}}) ({{.PackageDomainName}}, error) {
	logger := s.logger.StartTrace(ctx, "create")
   	ctx = logger.Context()
   	defer logger.FinishTrace()

	result, err := s.datastore.{{.NamesRepoCamel}}.Create(ctx, {{.NameLowerCamel}})
	if err != nil {
		logger.Error(err)
		return {{.PackageDomainName}}{}, errs.Internal{}
	}

	return result, nil
}

// Get - method for getting only one {{.NameLowerSpace}}.
func (s Service) Get(ctx context.Context, params filter.Filter) ({{.PackageDomainName}}, error) {
    logger := s.logger.StartTrace(ctx, "get")
    ctx = logger.Context()
    defer logger.FinishTrace()

	result, err := s.datastore.{{.NamesRepoCamel}}.Get(ctx, params)
	if err != nil {
		if errors.As(err, &repository.ErrNotFound{}) {
			logger.Debug(err)
			return {{.PackageDomainName}}{}, errs.NotFound{What: "{{.NameLowerCamel}}"}
		}

		logger.Error(err)

		return {{.PackageDomainName}}{}, errs.Internal{}
	}

	return result, nil
}

// GetAll - method for getting all {{.NamesLowerSpace}}.
func (s Service) GetAll(ctx context.Context, params filter.Filter) ({{.PackageDomainByPagination}}, error) {
    logger := s.logger.StartTrace(ctx, "get all")
    ctx = logger.Context()
    defer logger.FinishTrace()

	result, err := s.datastore.{{.NamesRepoCamel}}.GetAll(ctx, params)
	if err != nil {
		if errors.As(err, &repository.ErrNotFound{}) {
			logger.Debug(err)
			return {{.PackageDomainByPagination}}{}, errs.Empty{}
		}

		logger.Error(err)

		return {{.PackageDomainByPagination}}{}, errs.Internal{}
	}

	return result, nil
}

// Update - method for updating {{.NameLowerSpace}}.
func (s Service) Update(ctx context.Context, {{.NameLowerCamel}} {{.PackageDomainName}}) error {
    logger := s.logger.StartTrace(ctx, "update")
    ctx = logger.Context()
    defer logger.FinishTrace()

    if err := s.datastore.{{.NamesRepoCamel}}.Update(ctx, {{.NameLowerCamel}}); err != nil {
		logger.Error(err)
		return errs.Internal{}
	}

	return nil
}

// Delete - method for deleting {{.NameLowerSpace}}.
func (s Service) Delete(ctx context.Context, {{ .FieldIDCamel }} {{ .FieldIDType }}) error {
    logger := s.logger.StartTrace(ctx, "delete")
    ctx = logger.Context()
    defer logger.FinishTrace()

	if err := s.datastore.{{.NamesRepoCamel}}.Delete(ctx, {{ .FieldIDCamel }}); err != nil {
		logger.Error(err)
		return errs.Internal{}
	}

	return nil
}
