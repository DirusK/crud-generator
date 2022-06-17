package models

import (
	"fmt"
	"strings"

	"crud-generator-gui/internal/models"
)

func dbNullType(fieldType models.Type) string {
	switch fieldType {
	case models.TypeString, models.TypeEnum:
		return "sql.NullString"
	case models.TypeInt, models.TypeInt8, models.TypeInt16, models.TypeInt32, models.TypeInt64:
		return "sql.NullInt64"
	case models.TypeUint, models.TypeUint8, models.TypeUint16, models.TypeUint32, models.TypeUint64:
		return "sql.NullInt64"
	case models.TypeBool:
		return "sql.NullBool"
	case models.TypeTime:
		return "sql.NullTime"
	case models.TypeDecimal:
		return "decimal.NullDecimal"
	case models.TypeByte:
		return "sql.NullByte"
	case models.TypeFloat64, models.TypeFloat32:
		return "sql.NullFloat64"
	case models.TypeUUID:
		return "uuid.NullUUID"
	default:
		return ""
	}
}

func (e Entity) DatabaseModel() string {
	var result []string

	for _, field := range e.Fields {
		tag := fmt.Sprintf(`db:"%s"`, field.NameSnake())

		var fieldType string

		switch field.Nullable {
		case true:
			fieldType = dbNullType(field.Type)
		case false:
			if field.Type == models.TypeEnum {
				fieldType = e.PackageLower() + "." + field.NameCamel(true)
			} else {
				fieldType = field.Type.String()
			}
		}

		result = append(result, fmt.Sprintf("%s %s `%s`", field.NameCamel(true), fieldType, tag))
	}

	return strings.Join(result, "\n\t\t")
}

func (e Entity) DomainModel() string {
	var result []string

	for _, field := range e.Fields {
		var (
			fieldType string
			fieldName = field.NameCamel(true)
		)

		if field.Type == models.TypeEnum {
			fieldType = fieldName
		} else {
			fieldType = field.Type.String()
		}

		result = append(result, fieldName+" "+fieldType)
	}

	return strings.Join(result, "\n\t\t")
}
