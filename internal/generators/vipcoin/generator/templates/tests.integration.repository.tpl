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
	"github.com/olekukonko/tablewriter"
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
			name: "[success] create {{.NameLowerCamel}}",
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
			name: "[success] get one {{.NameLowerCamel}}",
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
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader([]string{"{{.NameCamel}}"})
					table.SetRowLine(true)
					table.Append([]string{fmt.Sprintf("%+v", got)})
					table.Render()
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
			name: "[success] get all {{.NamesLowerCamel}}",
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
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader([]string{"Index", "{{.NamesCamel}}"})
					table.SetRowLine(true)

					{{ if .WithPaginationCheck }}

					for idx, {{.NameLowerCamel}} := range got.{{.NamesCamel}} {
						table.Append([]string{strconv.Itoa(idx), fmt.Sprintf("%+v", {{.NameLowerCamel}})})
					}

					{{ else }}

					for idx, {{.NameLowerCamel}} := range got {
						table.Append([]string{strconv.Itoa(idx), fmt.Sprintf("%+v", {{.NameLowerCamel}})})
					}

					{{ end }}
					table.Render()
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
			name: "[success] delete {{.NameLowerCamel}}",
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
			name: "[success] update {{.NamesLowerCamel}}}",
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
