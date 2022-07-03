package models

import (
	"strings"

	"github.com/iancoleman/strcase"

	"crud-generator/internal/generators/vipcoin/helpers"
	"crud-generator/internal/models"
)

type (
	// Field for VipCoin generator
	Field struct {
		models.Field
	}
)

// NewField constructor.
func NewField(field models.Field) Field {
	return Field{Field: field}
}

func (f Field) IsEnum() bool {
	return f.Type == models.TypeEnum
}

func (f Field) GoFileSnakeWithExtension() string {
	return strcase.ToSnake(f.Name) + ".go"
}

func (f Field) NameCamel(withAbbreviation bool) string {
	result := strcase.ToCamel(f.Name)

	if withAbbreviation {
		return helpers.ReplaceAbbreviations(result)
	}

	return result
}

func (f Field) NameLowerCamel(withAbbreviation bool) string {
	result := strcase.ToLowerCamel(f.Name)

	if withAbbreviation {
		return helpers.ReplaceAbbreviations(result)
	}

	return result
}

func (f Field) NameSnake() string {
	return strcase.ToSnake(f.Name)
}

func (f Field) Reference() string {
	return strings.ToLower(f.Name[:1])
}
