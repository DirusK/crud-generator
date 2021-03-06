package app

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"crud-generator/pkg/printer"
)

// settings returns settings window.
func (a *App) settingsWindow() fyne.CanvasObject {
	emptyLine := canvas.NewText("", color.White)

	typeSelector := widget.NewSelect(generatorsString, func(s string) {
		a.generatorType = GeneratorType(s)
	})

	typeSelector.PlaceHolder = GeneratorVipCoin.String()
	a.generatorType = GeneratorVipCoin

	str := binding.NewString()

	moduleNameEntry := widget.NewEntry()
	moduleNameEntry.PlaceHolder = "Name"
	moduleNameEntry.OnChanged = func(module string) {
		a.settings.ModuleName = module
	}

	projectPathEntry := widget.NewEntryWithData(str)
	projectPathEntry.PlaceHolder = "Path"
	projectPathEntry.OnChanged = func(path string) {
		a.settings.ProjectPath = path

		parsed := strings.FieldsFunc(path, split)
		if len(parsed) > 0 {
			result := parsed[len(parsed)-1]
			moduleNameEntry.Text = result
			a.settings.ModuleName = result

			moduleNameEntry.Refresh()
		}
	}

	return container.New(
		layout.NewFormLayout(),
		layout.NewSpacer(),
		newSettingsTitle("Project"),
		emptyLine,
		emptyLine,
		newText("Project path", 16, color.Black, true, false),
		projectPathEntry,
		layout.NewSpacer(),
		prepareProjectPathButton(a, str),
		emptyLine,
		emptyLine,
		newText("Module name", 16, color.Black, true, false),
		moduleNameEntry,
		emptyLine,
		emptyLine,
		layout.NewSpacer(),
		newSettingsTitle("Generator"),
		emptyLine,
		emptyLine,
		newText("Type", 16, color.Black, true, false),
		typeSelector,
		emptyLine,
		emptyLine,
		layout.NewSpacer(),
		prepareCheckObjects(a),
	)
}

// newSettingsTitle returns container that stores title for settings.
func newSettingsTitle(text string) fyne.CanvasObject {
	return container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			newText(text, 20, color.Black, true, true),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}

// prepareCheckObjects return check boxes for generate options.
func prepareCheckObjects(a *App) fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 6)

	objects = append(objects, widget.NewCheck("Generate migration", func(selected bool) {
		a.settings.GenerateMigration = selected
	}))

	objects = append(objects, widget.NewCheck("Generate integration tests", func(selected bool) {
		a.settings.GenerateIntegrationTests = selected
	}))

	objects = append(objects, widget.NewCheck("Generate repository", func(selected bool) {
		a.settings.GenerateRepository = selected
	}))

	objects = append(objects, widget.NewCheck("Generate domain", func(selected bool) {
		a.settings.GenerateDomain = selected
	}))

	objects = append(objects, widget.NewCheck("Generate service", func(selected bool) {
		a.settings.GenerateService = selected
	}))

	objects = append(objects, widget.NewCheck("Generate handler", func(selected bool) {
		a.settings.GenerateHandler = selected
	}))

	return container.NewHBox(layout.NewSpacer(), container.NewVBox(objects...), layout.NewSpacer())
}

// prepareProjectPathButton return button for choosing project.
func prepareProjectPathButton(a *App, str binding.String) fyne.CanvasObject {
	return widget.NewButton("Open project", func() {
		folderOpen := dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
			if list == nil {
				return
			}

			if err = str.Set(list.Path()); err != nil {
				printer.Error(TagGUI, err, "can't set project path")
				dialog.ShowError(err, a.window)
				return
			}

			a.settings.ProjectPath = list.Path()
		}, a.window)

		folderOpen.Resize(fyne.NewSize(700, 500))
		folderOpen.Show()
	})
}
