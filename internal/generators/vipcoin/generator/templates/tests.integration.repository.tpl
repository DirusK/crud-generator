{{- /*gotype: crud-generator/internal/generators/vipcoin/models.Entity*/ -}}
{{ .Copyright }}

package {{.PackageLower}}

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"git.ooo.ua/vipcoin/lib/filter"
	"github.com/stretchr/testify/assert"

	"{{.ModuleNameLower}}/_tests/integration/repository"
	domain "{{.ModuleNameLower}}/internal/api/domain/{{.PackageLower}}"
)

func TestRepository_Create(t *testing.T) {
	type args struct {
		ctx    context.Context
		entity domain.{{.NameCamel}}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] create {{.NameLowerSpace}}",
			args: args{
				ctx: context.Background(),
				entity: domain.Fake{{.NameCamel}}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.Datastore.{{.NamesRepoCamel}}.Create(tt.args.ctx, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if assert.NotEmpty(t, got) {
					t.Logf("got {{.NameLowerCamel}}: %+v", got)
				}
			}
		})
	}
}

func TestRepository_Get(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] get one {{.NameLowerSpace}}",
			args: args{
				ctx:    context.Background(),
				filter: filter.NewFilter(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.Datastore.{{.NamesRepoCamel}}.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if assert.NotEmpty(t, got) {
					t.Logf("got {{.NameLowerCamel}}: %+v", got)
				}
			}
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] get all {{.NamesLowerSpace}}",
			args: args{
				ctx:    context.Background(),
				filter: filter.NewFilter().SetLimit(2),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.Datastore.{{.NamesRepoCamel}}.GetAll(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if assert.NotEmpty(t, got) {
					{{ if .WithPaginationCheck }}
					for idx, {{.NameLowerCamel}} := range got.{{.NamesCamel}} {
						t.Logf("index: %d, {{.NameLowerCamel}}: %+v", idx, {{.NameLowerCamel}})
					}
					{{ else }}
					for idx, {{.NameLowerCamel}} := range got {
						t.Logf("index: %d, {{.NameLowerCamel}}: %+v", idx, {{.NameLowerCamel}})
					}
					{{ end }}
				}
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		{{.FieldIDCamel}}  {{.FieldIDType}}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] delete {{.NameLowerSpace}}",
			args: args{
				ctx: context.Background(),
				{{.FieldIDCamel}}:  {{ if eq `uuid.UUID` .FieldIDType}} uuid.MustParse("set your uuid here") {{ else }} 1 {{end}}, // TODO: Replace ID value.
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.Datastore.{{.NamesRepoCamel}}.Delete(tt.args.ctx, tt.args.{{.FieldIDCamel}})
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {
	type args struct {
		ctx    context.Context
		{{.NamesLowerCamel}} []domain.{{.NameCamel}}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] update {{.NamesLowerSpace}}}",
			args: args{
				ctx: context.Background(),
				{{.NamesLowerCamel}}: []domain.{{.NameCamel}}{ // TODO: Replace with your data.
					domain.Fake{{.NameCamel}}(domain.FakeWith{{.FieldIDCamel}}({{ if eq `uuid.UUID` .FieldIDType}} uuid.MustParse("set your uuid here") {{ else }} 1 {{end}})),
					domain.Fake{{.NameCamel}}(domain.FakeWith{{.FieldIDCamel}}({{ if eq `uuid.UUID` .FieldIDType}} uuid.MustParse("set your uuid here") {{ else }} 2 {{end}})),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repository.Datastore.{{.NamesRepoCamel}}.Update(tt.args.ctx, tt.args.{{.NamesLowerCamel}}...); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
