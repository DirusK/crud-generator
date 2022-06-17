package models

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

type (
	// Field describes field of Entity.
	Field struct {
		Name       string
		Type       Type
		EnumValues []string
		Nullable   bool
	}
)

func (f Field) IsEnum() bool {
	return f.Type == TypeEnum
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

// Validate field.
func (f Field) Validate() error {
	var errs []string

	if f.Name == "" {
		errs = append(errs, fmt.Sprintf("name is required for field <%s>", f.Name))
	}

	if f.Type == "" {
		errs = append(errs, fmt.Sprintf("type is required for field <%s>", f.Name))
	}

	if f.Type == TypeEnum {
		if len(f.EnumValues) == 0 {
			errs = append(errs, fmt.Sprintf("enum values is required for field <%s>", f.Name))
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}
