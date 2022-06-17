package app

import (
	"fyne.io/fyne/v2"

	"crud-generator-gui/internal/generators"
	"crud-generator-gui/internal/models"
)

type (
	// App main application structure.
	App struct {
		meta          Meta
		window        fyne.Window
		generatorType GeneratorType
		generator     generators.Generator
		entity        models.Entity
		settings      models.Settings
	}

	// Meta stores all meta information for the application.
	Meta struct {
		AppName string
	}
)

// New application constructor.
func New(appName string) *App {
	return &App{
		meta:          Meta{AppName: appName},
		generatorType: GeneratorVipCoin,
	}
}

// Run stars the application.
func (a *App) Run() {
	a.mainWindow().ShowAndRun()
}

// setGenerator setups generator to the application structure.
func (a *App) setGenerator(init GeneratorInit) {
	a.generator = init(a.entity, a.settings)
}

// generate code by the specified generator.
func (a App) generate() error {
	var err error

	if a.settings.GenerateMigration {
		err = a.generator.GenerateMigration()
		if err != nil {
			return err
		}
	}

	if a.settings.GenerateRepository {
		err = a.generator.GenerateRepository()
		if err != nil {
			return err
		}
	}

	if a.settings.GenerateIntegrationTests {
		err = a.generator.GenerateIntegrationTests()
		if err != nil {
			return err
		}
	}

	if a.settings.GenerateDomain {
		err = a.generator.GenerateDomain()
		if err != nil {
			return err
		}
	}

	if a.settings.GenerateService {
		err = a.generator.GenerateService()
		if err != nil {
			return err
		}
	}

	if a.settings.GenerateHandler {
		err = a.generator.GenerateHandler()
		if err != nil {
			return err
		}
	}

	return nil
}
