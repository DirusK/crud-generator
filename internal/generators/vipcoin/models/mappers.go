package models

import (
	"fmt"
	"strings"

	"crud-generator-gui/internal/models"
)

func toNullConverter(fieldType models.Type) string {
	switch fieldType {
	case models.TypeString:
		return `repository.ToNullString(%s)`
	case models.TypeUUID:
		return `repository.ToNullUUID(%s)`
	case models.TypeInt64:
		return `repository.ToNullInt64(%s)`
	case models.TypeTime:
		return `repository.ToNullTime(%s)`
	case models.TypeBool:
		return `repository.ToNullBool(%s)`
	case models.TypeByte:
		return `repository.ToNullByte(%s)`
	case models.TypeDecimal:
		return `decimal.NewNullDecimal(%s)`
	case models.TypeInt, models.TypeInt8, models.TypeInt16, models.TypeInt32:
		return `repository.ToNullInt64(int64(%s))`
	case models.TypeFloat64:
		return `repository.ToNullFloat64(%s)`
	case models.TypeFloat32:
		return `repository.ToNullFloat64(float64(%s))`
	case models.TypeUint, models.TypeUint8, models.TypeUint16, models.TypeUint32, models.TypeUint64:
		return `repository.ToNullInt64(int64(%s))`
	case models.TypeEnum:
		return `repository.ToNullString(%s.String())`
	default:
		return ""
	}
}

func fromNullConverter(field Field, packageLower string) string {
	switch field.Type {
	case models.TypeString:
		return `%s.String`
	case models.TypeByte:
		return `%s.Byte`
	case models.TypeBool:
		return `%s.Bool`
	case models.TypeUUID:
		return `%s.UUID`
	case models.TypeTime:
		return `%s.Time`
	case models.TypeInt:
		return `int(%s.Int64)`
	case models.TypeInt8:
		return `int8(%s.Int64)`
	case models.TypeInt16:
		return `int16(%s.Int64)`
	case models.TypeInt32:
		return `int32(%s.Int64)`
	case models.TypeInt64:
		return `%s.Int64`
	case models.TypeFloat64:
		return `%s.Float64`
	case models.TypeFloat32:
		return `float32(%s.Float64)`
	case models.TypeDecimal:
		return `%s.Decimal`
	case models.TypeUint:
		return `uint(%s.Int64)`
	case models.TypeUint8:
		return `uint8(%s.Int64)`
	case models.TypeUint16:
		return `uint16(%s.Int64)`
	case models.TypeUint32:
		return `uint32(%s.Int64)`
	case models.TypeUint64:
		return `uint64(%s.Int64)`
	case models.TypeEnum:
		return packageLower + `.` + field.NameCamel(true) + `(%s.String)`
	default:
		return ""
	}
}

func (e Entity) FromDomainToDatabase() string {
	var result []string

	for _, field := range e.Fields {
		fieldName := field.NameCamel(true)

		domain := "domain." + fieldName
		if field.Nullable {
			domain = fmt.Sprintf(toNullConverter(field.Type), domain)
		}

		result = append(result, fmt.Sprintf("%s: %s,", fieldName, domain))
	}

	return strings.Join(result, "\n\t\t")
}

func (e Entity) FromDatabaseToDomain() string {
	var result []string

	for _, field := range e.Fields {
		fieldName := field.NameCamel(true)

		db := e.Reference() + "." + fieldName
		if field.Nullable {
			db = fmt.Sprintf(fromNullConverter(field, e.PackageLower()), db)
		}

		result = append(result, fmt.Sprintf("%s: %s,", fieldName, db))
	}

	return strings.Join(result, "\n\t\t")
}

func (e Entity) FromDomainToResponse() string {
	var result []string

	for _, field := range e.Fields {
		name := field.NameCamel(true)
		result = append(result, fmt.Sprintf("%s: domain.%s,", name, name))
	}

	return strings.Join(result, "\n\t\t")
}

func (e Entity) FromRequestToDomain() string {
	var result []string

	for _, field := range e.Fields {
		name := field.NameCamel(true)
		result = append(result, fmt.Sprintf("%s: %s.%s,", name, e.Reference(), name))
	}

	return strings.Join(result, "\n\t\t")
}
