package models

import (
	"strings"

	"github.com/pkg/errors"
)

// Settings store user custom selected settings for the application.
type Settings struct {
	ProjectPath string
	ModuleName  string

	GenerateMigration        bool
	GenerateIntegrationTests bool
	GenerateRepository       bool
	GenerateDomain           bool
	GenerateService          bool
	GenerateHandler          bool
}

// Validate settings structure.
func (s Settings) Validate() error {
	var errs []string

	if s.ProjectPath == "" {
		errs = append(errs, "project path is required")
	}

	if s.ModuleName == "" {
		errs = append(errs, "project name is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}
