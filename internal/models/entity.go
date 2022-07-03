package models

import (
	"strings"

	"github.com/pkg/errors"
)

type (
	// Entity describes entity to be generated.
	Entity struct {
		Name           string
		Copyright      string
		Package        string
		Table          string
		WithPagination bool
		Fields         []Field
	}
)

// Validate entity with fields.
func (e Entity) Validate() error {
	var errs []string

	if e.Name == "" {
		errs = append(errs, "entity name is required")
	}
	if e.Package == "" {
		errs = append(errs, "package name is required")
	}
	if e.Table == "" {
		errs = append(errs, "table name is required")
	}

	if len(e.Fields) == 0 {
		errs = append(errs, "entity must not be empty")
	}

	for idx := range e.Fields {
		if err := e.Fields[idx].Validate(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}
