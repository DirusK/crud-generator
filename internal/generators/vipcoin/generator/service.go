package generator

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func (g Generator) GenerateService() error {
	dirPath := fmt.Sprintf("%s/internal/api/services/%s/", g.Settings.ProjectPath, g.Entity.PackageLower())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "can't make service directory")
	}

	if err := g.generateServiceMethods(dirPath); err != nil {
		return err
	}

	return nil
}

func (g Generator) generateServiceMethods(dirPath string) error {
	return g.executeTemplate(
		"service.methods.tpl",
		dirPath+g.Entity.GoFileSnakeWithExtension(),
		true,
		g.Entity,
	)
}
