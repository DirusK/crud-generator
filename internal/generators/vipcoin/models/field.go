package models

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"crud-generator-gui/internal/models"
)

type (
	// Field for VipCoin generator
	Field struct {
		models.Field
	}
)

func (f Field) IsEnum() bool {
	return f.Type == models.TypeEnum
}

func (f Field) GoFileSnakeWithExtension() string {
	return strcase.ToSnake(f.Name) + ".go"
}

func (f Field) NameCamel(withAbbreviation bool) string {
	result := strcase.ToCamel(f.Name)

	if withAbbreviation {
		return replaceAbbreviations(result)
	}

	return result
}

func (f Field) NameLowerCamel(withAbbreviation bool) string {
	result := strcase.ToLowerCamel(f.Name)

	if withAbbreviation {
		return replaceAbbreviations(result)
	}

	return result
}

func (f Field) NameSnake() string {
	return strcase.ToSnake(f.Name)
}

func (f Field) Reference() string {
	return strings.ToLower(f.Name[:1])
}

func (f Field) EnumMap() string {
	var constants string
	for _, value := range f.EnumValues {
		constants += fmt.Sprintf("%s: {},\n\t", f.NameCamel(true)+strcase.ToCamel(value))
	}

	return constants
}

func (f Field) EnumArray() string {
	var fields string
	for _, value := range f.EnumValues {
		fields += f.NameCamel(true) + strcase.ToCamel(value) + ".String(), \n"
	}

	return fields
}

func (f Field) EnumConstants() string {
	var constants []string
	for _, value := range f.EnumValues {
		nameCamel := f.NameCamel(true)
		enum := nameCamel + strcase.ToCamel(value)

		constants = append(
			constants,
			fmt.Sprintf(`%s %s = "%s"`, enum, nameCamel, value),
		)
	}

	return strings.Join(constants, "\n\t")
}

func (f Field) MigrationEnumArray() string {
	var result []string
	for _, value := range f.EnumValues {
		result = append(result, fmt.Sprintf(`'%s'`, value))
	}

	return strings.Join(result, ",")
}
