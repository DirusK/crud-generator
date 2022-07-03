package generator

import (
	"os"

	"github.com/pkg/errors"
)

func (g Generator) GenerateDomain() error {
	dirPath := g.Settings.ProjectPath + "/internal/api/domain/" + g.Entity.PackageLower() + "/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "can't make repository directory")
	}

	if err := g.generateDomainModel(dirPath); err != nil {
		return err
	}

	if err := g.generateDomainEnums(dirPath); err != nil {
		return err
	}

	if err := g.generateFake(dirPath); err != nil {
		return err
	}

	return nil
}

func (g Generator) generateDomainModel(dirPath string) error {
	return g.executeTemplate(
		"./internal/generators/vipcoin/templates/domain.model.tpl",
		dirPath+g.Entity.GoFileSnakeWithExtension(),
		true,
		g.Entity,
	)
}

func (g Generator) generateDomainEnums(dirPath string) error {
	for _, enum := range g.Entity.GetEnumFields() {
		if err := g.executeTemplate(
			"./internal/generators/vipcoin/templates/domain.enum.tpl",
			dirPath+enum.GoFileSnakeWithExtension(),
			true,
			struct {
				Copyright      string
				PackageLower   string
				EnumCamel      string
				EnumLowerCamel string
				EnumsCamel     string
				NameCamel      string
				EnumConstants  string
				EnumMap        string
				Reference      string
			}{
				Copyright:      g.Entity.Copyright,
				PackageLower:   g.Entity.PackageLower(),
				EnumCamel:      enum.NameCamel(true),
				EnumLowerCamel: enum.NameLowerCamel(true),
				NameCamel:      g.Entity.NameCamel(),
				EnumConstants:  enum.EnumConstants(),
				EnumMap:        enum.EnumMap(),
				Reference:      enum.Reference(),
			},
		); err != nil {
			return err
		}
	}

	return nil
}

func (g Generator) generateFake(dirPath string) error {
	return g.executeTemplate(
		"./internal/generators/vipcoin/templates/domain.fake.tpl",
		dirPath+"fake.go",
		true,
		g.Entity,
	)
}
