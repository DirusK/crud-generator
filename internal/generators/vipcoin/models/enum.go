package models

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

const IsEnumCheckTemplate = `// Is{{.NameWithValueCamel}} check if {{.EntityNameCamel}} {{.NameLowerCamel}} is {{.NameWithValueCamel}}.
func ({{.Reference}} {{.EntityNameCamel}}) Is{{.NameWithValueCamel}}() bool { return {{.Reference}}.{{.NameCamel}} == {{.NameWithValueCamel}} }`

func (e Entity) FieldsEnum() string {
	var constants []string
	for _, field := range e.Fields {
		constants = append(
			constants,
			fmt.Sprintf(`Field%s = "%s"`, field.NameCamel(true), field.NameSnake()),
		)
	}

	return strings.Join(constants, "\n\t")
}

func (e Entity) GetEnumFields() []Field {
	var enums []Field

	for _, field := range e.Fields {
		if field.IsEnum() {
			enums = append(enums, field)
		}
	}

	return enums
}

func (e Entity) IsEnumCheck() string {
	var result []string

	for _, field := range e.GetEnumFields() {
		for _, value := range field.EnumValues {
			fieldNameCamel := field.NameCamel(true)

			result = append(result, executeTemplateFromString(IsEnumCheckTemplate, struct {
				NameWithValueCamel string
				EntityNameCamel    string
				NameLowerCamel     string
				Reference          string
				NameCamel          string
			}{
				NameWithValueCamel: fieldNameCamel + strcase.ToCamel(value),
				EntityNameCamel:    e.NameCamel(),
				NameLowerCamel:     field.NameLowerCamel(true),
				Reference:          e.Reference(),
				NameCamel:          fieldNameCamel,
			}))
		}
	}

	return strings.Join(result, "\n\n")
}
