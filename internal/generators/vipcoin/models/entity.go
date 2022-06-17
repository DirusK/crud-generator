package models

import (
	"strings"

	"github.com/iancoleman/strcase"

	"crud-generator-gui/internal/models"
)

type (
	// Entity for VipCoin generator
	Entity struct {
		ModuleName          string
		Name                string
		Package             string
		Table               string
		MigrationExtensions migrationExtensions
		WithPagination      bool
		Fields              []Field
	}

	migrationExtensions map[models.Type]struct{}
)

func NewEntity(moduleName, name, packageName, tableName string, withPagination bool, fields []models.Field) Entity {
	var (
		f   []Field
		ext = make(migrationExtensions)
	)

	for _, field := range fields {
		switch field.Type {
		case models.TypeUUID:
			ext[models.TypeUUID] = struct{}{}
		}

		f = append(f, Field{Field: field})
	}

	return Entity{
		ModuleName:          moduleName,
		Name:                name,
		Package:             packageName,
		Table:               tableName,
		MigrationExtensions: ext,
		WithPagination:      withPagination,
		Fields:              f,
	}
}

func (e Entity) WithPaginationCheck() bool {
	return e.WithPagination
}

func (e Entity) PackageLower() string {
	return strings.ToLower(e.Package)
}

func (e Entity) NamesKebab() string {
	return strcase.ToKebab(e.Name + "s")
}

func (e Entity) ModuleNameLower() string {
	return strings.ToLower(e.ModuleName)
}

func (e Entity) NameLowerCamel() string {
	return strcase.ToLowerCamel(e.Name)
}

func (e Entity) NameCamel() string {
	return strcase.ToCamel(e.Name)
}

func (e Entity) NameSnake() string {
	return strcase.ToSnake(e.Name)
}

func (e Entity) NamesLowerCamel() string {
	return strcase.ToLowerCamel(e.Name) + "s"
}

func (e Entity) NamesCamel() string {
	return strcase.ToCamel(e.Name) + "s"
}

func (e Entity) ListCamel() string {
	return strcase.ToCamel(e.Name + "List")
}

func (e Entity) ListLowerCamel() string {
	return strcase.ToLowerCamel(e.Name + "List")
}

func (e Entity) NamesRepoCamel() string {
	return strcase.ToCamel(e.Name + "sRepo")
}

func (e Entity) NamesServiceCamel() string {
	return strcase.ToCamel(e.Name + "sService")
}

func (e Entity) Reference() string {
	return strings.ToLower(e.Name[:1])
}

func (e Entity) GoFileSnakeWithExtension() string {
	return strcase.ToSnake(e.Name+"s") + ".go"
}

func (e Entity) GoFileSnakeWithoutExtension() string {
	return strcase.ToSnake(e.Name + "s")
}

func (e Entity) Interface() string {
	return strcase.ToCamel(e.Name + "s")
}

func (e Entity) PackageDomainByPagination() string {
	if e.WithPagination {
		return e.Package + "." + e.ListCamel()
	}

	return e.Package + "." + e.NamesCamel()
}

func (e Entity) PackageDomainName() string {
	return e.Package + "." + e.NameCamel()
}

func (e Entity) PackageDomainNameList() string {
	return e.Package + "." + e.NameCamel() + "List"
}

func (e Entity) PackageDomainNames() string {
	return e.Package + "." + e.NamesCamel()
}

func (e Entity) FieldIDSnake() string {
	return strcase.ToSnake(e.Fields[0].Name)
}

func (e Entity) FieldIDCamel() string {
	return e.Fields[0].NameCamel(true)
}

func (e Entity) FieldIDType() string {
	return e.Fields[0].Type.String()
}
