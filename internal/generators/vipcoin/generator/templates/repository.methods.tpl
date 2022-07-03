{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"git.ooo.ua/vipcoin/lib/filter"

	"{{.ModuleName}}/internal/api/domain/{{.PackageLower}}"
	"{{.ModuleName}}/internal/api/repository"
)

//go:generate ifacemaker -f {{.GoFileSnakeWithExtension}} -s Repository -p repository -i {{.Interface}} -y "{{.Interface}} - describe an interface for working with {{.NamesLowerCamel}} database models."

var _ repository.{{.Interface}} = &Repository{}

// Repository implements repository.{{.Interface}}
type Repository struct {
	queryTimeout time.Duration
	db           repository.DatabaseExecutor
}

// NewRepository constructor.
func NewRepository(qt time.Duration, db repository.DatabaseExecutor) *Repository {
	return &Repository{
		queryTimeout: qt,
		db:           db,
	}
}

// Create new {{.NameLowerCamel}} in database.
func (r Repository) Create(ctx context.Context, entity {{.PackageDomainName}}) ({{.PackageDomainName}}, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

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
		return {{.PackageDomainName}}{}, repository.ErrExecute{Cause: err.Error()}
	}

	defer func() {
		_ = stmt.Close()
	}()

	var result {{.NameLowerCamel}}
	if err = stmt.GetContext(ctx, &result, toDatabase(entity)); err != nil {
		return {{.PackageDomainName}}{}, repository.ErrExecute{Cause: err.Error()}
	}

	return result.toDomain(), nil
}

// Get one {{.NameLowerCamel}} from database by filter.
func (r Repository) Get(ctx context.Context, filter filter.Filter) ({{.PackageDomainName}}, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query, args := filter.SetLimit(1).Build(tableName)

	var result {{.NameLowerCamel}}

	if err := r.db.GetContext(ctx, &result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return {{.PackageDomainName}}{}, repository.ErrNotFound{What: "{{.NameLowerCamel}}"}
		}

		return {{.PackageDomainName}}{}, repository.ErrExecute{Cause: err.Error()}
	}

	return result.toDomain(), nil
}

{{ if .WithPaginationCheck }}

// GetAll {{.NamesLowerCamel}} from database by filter.
func (r Repository) GetAll(ctx context.Context, filter filter.Filter) ({{.PackageDomainNameList}}, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query, args := filter.
		SetLimit(0).
		SetOffset(0).
		SetSortMap(nil).
		Build(tableName, "count(*)")

	var total uint64
	if err := r.db.GetContext(ctx, &total, query, args...); err != nil {
		return {{.PackageDomainNameList}}{}, repository.ErrExecute{Cause: err.Error()}
	}

	if total == 0 {
		return {{.PackageDomainNameList}}{}, repository.ErrNotFound{What: "{{.NamesLowerCamel}}"}
	}

	query, args = filter.Build(tableName)

	var result {{.ListLowerCamel}}
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return {{.PackageDomainNameList}}{}, repository.ErrExecute{Cause: err.Error()}
	}

	return result.toDomain(total), nil
}

{{ else }}

// GetAll {{.NamesLowerCamel}} from database by filter.
func (r Repository) GetAll(ctx context.Context, filter filter.Filter) ({{ .PackageDomainNames }}, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query, args := filter.Build(tableName)

	var result {{.ListLowerCamel}}
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return {{ .PackageDomainNames }}{}, repository.ErrExecute{Cause: err.Error()}
	}

	if len(result) == 0 {
		return {{ .PackageDomainNames }}{}, repository.ErrNotFound{What: "{{.NamesLowerCamel}}"}
	}

	return result.toDomain(), nil
}

{{ end }}

// Update {{.NameLowerCamel}} in database.
func (r Repository) Update(ctx context.Context, {{.NamesLowerCamel}} ...{{.PackageDomainName}}) error {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `
		UPDATE {{.TableName}} SET
			{{.UpdateQuery}}
		WHERE {{.FieldIDSnake}} = :{{.FieldIDSnake}}
		`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return repository.ErrExecute{Cause: err.Error()}
	}

	for idx := range {{.NamesLowerCamel}} {
		if _, err = stmt.ExecContext(ctx, toDatabase({{.NamesLowerCamel}}[idx])); err != nil {
			return repository.ErrExecute{Cause: err.Error()}
		}
	}

	return nil
}

// Delete {{.NameLowerCamel}} in database.
func (r Repository) Delete(ctx context.Context, {{.FieldIDCamel}} {{.FieldIDType}}) error {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := ` DELETE FROM {{.TableName}} WHERE {{.FieldIDSnake}} = $1 `

	if _, err := r.db.ExecContext(ctx, query, ID); err != nil {
		return repository.ErrExecute{Cause: err.Error()}
	}

	return nil
}