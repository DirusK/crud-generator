package models

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type (
	// Field describes field of Entity.
	Field struct {
		Name       string
		Type       Type
		EnumValues []string
		Default    string
		Check      string
		Validation string
		References string
		Nullable   bool
		Omitempty  bool
		Unique     bool
	}
)

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
