package generator

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func (g Generator) GenerateHandler() error {
	dirPath := fmt.Sprintf("%s/internal/api/delivery/http/%s/", g.Settings.ProjectPath, g.Entity.PackageLower())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "can't make handler directory")
	}

	if err := g.generateHandlerMethods(dirPath); err != nil {
		return err
	}

	if err := g.generateHandlerRequestModels(dirPath); err != nil {
		return err
	}

	if err := g.generateHandlerResponseModels(dirPath); err != nil {
		return err
	}
	return nil
}

func (g Generator) generateHandlerMethods(dirPath string) error {
	return g.executeTemplate(
		"./internal/generators/vipcoin/templates/handler.methods.tpl",
		dirPath+g.Entity.GoFileSnakeWithExtension(),
		true,
		g.Entity,
	)
}

func (g Generator) generateHandlerRequestModels(dirPath string) error {
	return g.executeTemplate(
		"./internal/generators/vipcoin/templates/handler.request.tpl",
		dirPath+"request.go",
		true,
		g.Entity,
	)
}

func (g Generator) generateHandlerResponseModels(dirPath string) error {
	return g.executeTemplate(
		"./internal/generators/vipcoin/templates/handler.response.tpl",
		dirPath+"response.go",
		true,
		g.Entity,
	)
}
