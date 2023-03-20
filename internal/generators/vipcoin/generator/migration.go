package generator

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func (g Generator) GenerateMigration() error {
	dirPath := g.Settings.ProjectPath + "/cmd/dbschema/migrations"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "can't make repository directory")
	}

	return g.executeTemplate(
		"migration.tpl",
		fmt.Sprintf("%s/%s-%s.sql", dirPath, currentTimeForMigration(), g.Entity.TableName()),
		false,
		g.Entity,
	)
}
