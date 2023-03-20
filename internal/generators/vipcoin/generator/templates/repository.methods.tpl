{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

import (
	"context"
	"database/sql"
	"errors"
	"time"

    "git.ooo.ua/vipcoin/lib/log"
    "git.ooo.ua/vipcoin/lib/errs"
	"git.ooo.ua/vipcoin/lib/filter"

	"{{.ModuleName}}/internal/api/domain/{{.PackageLower}}"
	"{{.ModuleName}}/internal/api/repository"
)

//go:generate ifacemaker -f {{.GoFileSnakeWithExtension}} -s Repository -p repository -i {{.Interface}} -y "{{.Interface}} - describe an interface for working with {{.NamesLowerCamel}} database models."

var _ repository.{{.Interface}} = &Repository{}

// Repository implements repository.{{.Interface}}
type Repository struct {
	db           database.Executor
	logger       log.Logger
}

// NewRepository constructor.
func NewRepository(db database.Executor, logger log.Logger) *Repository {
	return &Repository{
		db:           db,
		logger:       logger.With("{{.NamesLowerSpace}} repository"),
	}
}

// Create new {{.NameLowerCamel}} in database.
func (r Repository) Create(ctx context.Context, entity {{.PackageDomainName}}) ({{.PackageDomainName}}, error) {
	logger := r.logger.StartTrace(ctx, "create")
	ctx = logger.Context()
	defer logger.FinishTrace()

	query := `
		INSERT INTO {{.TableName}} (
			{{.InsertFields}}
		) VALUES (
			{{.InsertValues}}
		) RETURNING
			{{.SelectFields}},
			created_at,
			updated_at
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return {{.PackageDomainName}}{}, errs.Internal{Cause: err.Error()}
	}

	defer func() {
		_ = stmt.Close()
	}()

	var result {{.NameLowerCamel}}
	if err = r.db.NewStatement(stmt).GetContext(ctx, &result, toDatabase(entity)); err != nil {
		return {{.PackageDomainName}}{}, errs.Internal{Cause: err.Error()}
	}

	return result.toDomain(), nil
}

// Get one {{.NameLowerCamel}} from database by filter.
func (r Repository) Get(ctx context.Context, filter filter.Filter) ({{.PackageDomainName}}, error) {
	logger := r.logger.StartTrace(ctx, "get")
	ctx = logger.Context()
	defer logger.FinishTrace()

	query, args := filter.SetLimit(1).Build(tableName)

	var result {{.NameLowerCamel}}

	if err := r.db.GetContext(ctx, &result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return {{.PackageDomainName}}{}, errs.NotFound{What: "{{.NameSnake}}"}
		}

		return {{.PackageDomainName}}{}, errs.Internal{Cause: err.Error()}
	}

	return result.toDomain(), nil
}

{{ if .WithPaginationCheck }}

// GetAll {{.NamesLowerCamel}} from database by filter.
func (r Repository) GetAll(ctx context.Context, filter filter.Filter) ({{.PackageDomainNameList}}, error) {
	logger := r.logger.StartTrace(ctx, "get all")
	ctx = logger.Context()
	defer logger.FinishTrace()

	query, args := filter.
		SetLimit(0).
		SetOffset(0).
		SetSortMap(nil).
		Build(tableName, "count(*)")

	var total uint64
	if err := r.db.GetContext(ctx, &total, query, args...); err != nil {
		return {{.PackageDomainNameList}}{}, errs.Internal{Cause: err.Error()}
	}

	if total == 0 {
		return {{.PackageDomainNameList}}{}, errs.NotFound{What: "{{.NamesSnake}}"}
	}

	query, args = filter.Build(tableName)

	var result {{.ListLowerCamel}}
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return {{.PackageDomainNameList}}{}, errs.Internal{Cause: err.Error()}
	}

	return result.toDomain(total), nil
}

{{ else }}

// GetAll {{.NamesLowerCamel}} from database by filter.
func (r Repository) GetAll(ctx context.Context, filter filter.Filter) ({{ .PackageDomainNames }}, error) {
	logger := r.logger.StartTrace(ctx, "get all")
	ctx = logger.Context()
	defer logger.FinishTrace()

	query, args := filter.Build(tableName)

	var result {{.ListLowerCamel}}
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return {{ .PackageDomainNames }}{}, errs.Internal{Cause: err.Error()}
	}

	if len(result) == 0 {
		return {{ .PackageDomainNames }}{}, errs.NotFound{What: "{{.NamesSnake}}"}
	}

	return result.toDomain(), nil
}

{{ end }}

// Update {{.NameLowerCamel}} in database.
func (r Repository) Update(ctx context.Context, {{.NamesLowerCamel}} ...{{.PackageDomainName}}) error {
	logger := r.logger.StartTrace(ctx, "update")
	ctx = logger.Context()
	defer logger.FinishTrace()

	query := `
		UPDATE {{.TableName}} SET
			{{.UpdateQuery}}
		WHERE {{.FieldIDSnake}} = :{{.FieldIDSnake}}
		`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

    wrappedStmt := r.db.NewStatement(stmt)

	for idx := range {{.NamesLowerCamel}} {
		if _, err = wrappedStmt.ExecContext(ctx, toDatabase({{.NamesLowerCamel}}[idx])); err != nil {
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// Delete {{.NameLowerSpace}} in database.
func (r Repository) Delete(ctx context.Context, {{.FieldIDCamel}} {{.FieldIDType}}) error {
	logger := r.logger.StartTrace(ctx, "delete")
	ctx = logger.Context()
	defer logger.FinishTrace()

	query := ` DELETE FROM {{.TableName}} WHERE {{.FieldIDSnake}} = $1 `

	if _, err := r.db.ExecContext(ctx, query, ID); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}