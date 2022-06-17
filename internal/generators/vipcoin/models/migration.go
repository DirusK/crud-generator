package models

import (
	"fmt"
	"strings"

	"crud-generator-gui/internal/models"
)

func migrationDatabaseType(field Field, tableName string) string {
	switch field.Type {
	case models.TypeString:
		return "text"
	case models.TypeEnum:
		return tableName + "_" + field.NameSnake()
	case models.TypeInt8, models.TypeInt16, models.TypeByte:
		return "smallint"
	case models.TypeInt, models.TypeInt32, models.TypeUint8, models.TypeUint16:
		return "integer"
	case models.TypeUint, models.TypeUint32, models.TypeInt64, models.TypeUint64:
		return "bigint"
	case models.TypeBool:
		return "boolean"
	case models.TypeTime:
		return "timestamp"
	case models.TypeDecimal:
		return "numeric(32,16)"
	case models.TypeFloat64, models.TypeFloat32:
		return "numeric(32,16)"
	case models.TypeUUID:
		return "uuid"
	default:
		return ""
	}
}

type statement struct {
	Up   string
	Down string
}

func newStatement(up string, down string) statement {
	return statement{
		Up:   up,
		Down: down,
	}
}

var extensions = map[models.Type]statement{
	models.TypeUUID: newStatement(
		`create extension if not exists "uuid-ossp";`,
		`drop extension if exists "uuid-ossp";`,
	),
}

func (e Entity) MigrationCreateExtensions() string {
	var result []string

	for fieldType := range e.MigrationExtensions {
		result = append(result, extensions[fieldType].Up)
	}

	return strings.Join(result, "\n")
}

func (e Entity) MigrationDropExtensions() string {
	var result []string

	for fieldType := range e.MigrationExtensions {
		result = append(result, extensions[fieldType].Down)
	}

	return strings.Join(result, "\n")
}

func (e Entity) MigrationCreateTypes() string {
	var result []string

	for _, field := range e.Fields {
		if field.Type != models.TypeEnum {
			continue
		}

		result = append(result, fmt.Sprintf(
			`create type %s_%s as enum (%s);`,
			e.TableName(),
			field.NameSnake(),
			field.MigrationEnumArray(),
		))
	}

	return strings.Join(result, "\n")
}

func (e Entity) MigrationDropTypes() string {
	var result []string

	for _, field := range e.Fields {
		if field.Type != models.TypeEnum {
			continue
		}

		result = append(result, fmt.Sprintf(
			`drop type %s_%s;`,
			e.TableName(),
			field.NameSnake(),
		))
	}

	return strings.Join(result, "\n")
}

func (e Entity) MigrationTableFields() string {
	checkNull := func(field Field) string {
		if field.Nullable {
			return "null"
		}

		return "not null"
	}

	checkIDField := func(field Field) string {
		switch field.Type {
		case models.TypeUUID:
			return "uuid default uuid_generate_v4() primary key"
		default:
			return "serial primary key"
		}
	}

	var result []string
	for idx, field := range e.Fields {
		if idx == 0 {
			result = append(result, fmt.Sprintf(
				"%s %s %s,",
				field.NameSnake(),
				checkIDField(field),
				checkNull(field),
			))

			continue
		}

		result = append(result, fmt.Sprintf(
			"%s %s %s,",
			field.NameSnake(),
			migrationDatabaseType(field, e.TableName()),
			checkNull(field),
		))
	}

	return strings.Join(result, "\n")
}
