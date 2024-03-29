package generator

import (
	"os"

	"github.com/pkg/errors"
)

func (g Generator) GenerateIntegrationTests() error {
	dirPath := g.Settings.ProjectPath + "/_tests/integration/repository/" + g.Entity.PackageLower() + "/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "can't make repository directory")
	}

	return g.executeTemplate(
		"tests.integration.repository.tpl",
		dirPath+g.Entity.GoFileSnakeWithoutExtension()+"_test.go",
		true,
		g.Entity,
	)
}
