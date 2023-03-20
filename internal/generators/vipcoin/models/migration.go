package models

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"crud-generator/internal/models"
)

// migrateStatement stores up and down statements for database migration.
type migrateStatement struct {
	Up   string
	Down string
	Used bool
}

// newMigrateStatement returns a new migration statement.
func newMigrateStatement(up string, down string) migrateStatement {
	return migrateStatement{
		Up:   up,
		Down: down,
	}
}

// extensions stores all possible extensions for migration.
var extensions = map[models.Type]migrateStatement{
	models.TypeUUID: newMigrateStatement(
		`create extension if not exists "uuid-ossp";`,
		`drop extension if exists "uuid-ossp";`,
	),
}

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
	case models.TypeDecimal, models.TypeCoins:
		return "numeric"
	case models.TypeFloat64, models.TypeFloat32:
		return "numeric"
	case models.TypeUUID:
		return "uuid"
	default:
		return ""
	}
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
			field.EnumMigrationArray(),
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
	prepareNull := func(field Field) string {
		if field.Nullable {
			return "null"
		}

		return "not null"
	}

	prepareIDField := func(field Field) string {
		switch field.Type {
		case models.TypeUUID:
			return "uuid default uuid_generate_v4() primary key"
		default:
			return "serial primary key"
		}
	}

	prepareDefault := func(field Field) string {
		if field.Default == "" {
			return ""
		}

		switch field.Type {
		case models.TypeString, models.TypeEnum:
			return fmt.Sprintf("default '%s'", strcase.ToSnake(field.Default))
		default:
			return "default " + field.Default
		}
	}

	prepareCheck := func(field Field) string {
		if field.Check == "" {
			return ""
		}

		return fmt.Sprintf("check(%s)", field.Check)
	}

	prepareUnique := func(field Field) string {
		if field.Unique {
			return "unique"
		}

		return ""
	}

	prepareReferences := func(field Field) string {
		if field.References != "" {
			return fmt.Sprintf("references %s", field.References)
		}

		return ""
	}

	mergeOptions := func(str ...string) string {
		return strings.Join(str, " ") + ","
	}

	var result []string
	for idx, field := range e.Fields {
		if idx == 0 {
			result = append(result, fmt.Sprintf(
				"%s %s %s,",
				field.NameSnake(),
				prepareIDField(field),
				prepareNull(field),
			))

			continue
		}

		result = append(result,
			mergeOptions(
				field.NameSnake(),
				migrationDatabaseType(field, e.TableName()),
				prepareDefault(field),
				prepareCheck(field),
				prepareUnique(field),
				prepareReferences(field),
				prepareNull(field),
			))
	}

	return strings.Join(result, "\n\t")
}
