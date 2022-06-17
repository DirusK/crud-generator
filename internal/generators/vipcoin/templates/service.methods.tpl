{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
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
	logger    log.Logger
	datastore *repository.Datastore
}

// NewService - constructor.
func NewService(datastore *repository.Datastore, logger log.Logger) *Service {
	return &Service{
		datastore: datastore,
		logger:    logger,
	}
}

// Get - method for getting only one {{.NameLowerCamel}}.
func (s Service) Get(ctx context.Context, params filter.Filter) ({{.PackageDomainName}}, error) {
	result, err := s.datastore.{{.NamesRepoCamel}}.Get(ctx, params)
	if err != nil {
		if errors.As(err, &repository.ErrNotFound{}) {
			s.logger.Debug(err)
			return {{.PackageDomainName}}{}, errs.NotFound{What: "{{.NameLowerCamel}}"}
		}

		s.logger.Error(err)

		return {{.PackageDomainName}}{}, errs.Internal{}
	}

	return result, nil
}

// GetAll - method for getting all {{.NamesLowerCamel}}.
func (s Service) GetAll(ctx context.Context, params filter.Filter) ({{.PackageDomainByPagination}}, error) {
	result, err := s.datastore.{{.NamesRepoCamel}}.GetAll(ctx, params)
	if err != nil {
		if errors.As(err, &repository.ErrNotFound{}) {
			s.logger.Debug(err)
			return {{.PackageDomainByPagination}}{}, errs.Empty{}
		}

		s.logger.Error(err)

		return {{.PackageDomainByPagination}}{}, errs.Internal{}
	}

	return result, nil
}

// Create - method for creating {{.NameLowerCamel}}.
func (s Service) Create(ctx context.Context, {{.NameLowerCamel}} {{.PackageDomainName}}) ({{.PackageDomainName}}, error) {
	result, err := s.datastore.{{.NamesRepoCamel}}.Create(ctx, {{.NameLowerCamel}})
	if err != nil {
		s.logger.Error(err)
		return {{.PackageDomainName}}{}, errs.Internal{}
	}

	return result, nil
}

// Update - method for updating {{.NameLowerCamel}}.
func (s Service) Update(ctx context.Context, {{.NameLowerCamel}} {{.PackageDomainName}}) error {
	if err := s.datastore.{{.NamesRepoCamel}}.Update(ctx, {{.NameLowerCamel}}); err != nil {
		s.logger.Error(err)
		return errs.Internal{}
	}

	return nil
}

// Delete - method for deleting {{.NameLowerCamel}}.
func (s Service) Delete(ctx context.Context, {{ .FieldIDCamel }} {{ .FieldIDType }}) error {
	if err := s.datastore.{{.NamesRepoCamel}}.Delete(ctx, {{ .FieldIDCamel }}); err != nil {
		s.logger.Error(err)
		return errs.Internal{}
	}

	return nil
}
