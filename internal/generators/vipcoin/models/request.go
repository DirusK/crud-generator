package models

import (
	"fmt"
	"strings"

	"crud-generator-gui/internal/models"
)

func (e Entity) GetAllFilter() string {
	var result []string

	for _, field := range e.Fields {
		var fieldType string
		switch field.Type {
		case models.TypeEnum:
			fieldType = "[]" + e.PackageLower() + "." + field.NameCamel(true)
		case models.TypeBool:
			fieldType = field.Type.String()
		case models.TypeDecimal, models.TypeCoins:
			fieldType = "[]" + models.TypeFloat64.String()
		case models.TypeTime:
			continue
		default:
			fieldType = "[]" + field.Type.String()
		}

		queryTag := fmt.Sprintf(`query:"%s"`, field.NameSnake())
		result = append(result, fmt.Sprintf("%s %s `%s`", field.NameCamel(true), fieldType, queryTag))
	}

	return strings.Join(result, "\n\t\t")
}
