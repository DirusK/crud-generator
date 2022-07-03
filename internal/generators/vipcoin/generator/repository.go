package generator

import (
	"os"

	"github.com/pkg/errors"
)

func (g Generator) GenerateRepository() error {
	dirPath := g.Settings.ProjectPath + "/internal/api/repository/" + g.Entity.PackageLower() + "/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "can't make repository directory")
	}

	if err := g.generateRepositoryMethods(dirPath); err != nil {
		return err
	}

	if err := g.generateRepositoryModels(dirPath); err != nil {
		return err
	}

	return nil
}

func (g Generator) generateRepositoryMethods(dirPath string) error {
	return g.executeTemplate(
		"repository.methods.tpl",
		dirPath+g.Entity.GoFileSnakeWithExtension(),
		true,
		g.Entity,
	)
}

func (g Generator) generateRepositoryModels(dirPath string) error {
	return g.executeTemplate(
		"repository.models.tpl",
		dirPath+"models.go",
		true,
		g.Entity,
	)
}
