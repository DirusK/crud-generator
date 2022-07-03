package app

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/iancoleman/strcase"

	"crud-generator-gui/internal/models"
	"crud-generator-gui/pkg/printer"
)

// defaultFieldName returns default name for field in entity creation.
const defaultFieldName = "Field"

type (
	// selectedField contains information about list of fields.
	//  ID - selected item ID from the list
	// 	selected - flag if item is chosen.
	selectedField struct {
		ID       int
		selected bool
	}

	// fieldOptions contains widgets for configuring field.
	fieldOptions struct {
		nameEntry       *widget.Entry
		enumEntry       *widget.Entry
		defaultEntry    *widget.Entry
		checkEntry      *widget.Entry
		validationEntry *widget.Entry
		referencesEntry *widget.Entry
		typeSelect      *widget.Select
		omitemptyCheck  *widget.Check
		nullableCheck   *widget.Check
		uniqueCheck     *widget.Check
	}
)

// newFieldOption inits new field option widgets.
func newFieldOption(app *App, selectedField *selectedField) fieldOptions {
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "CallbackStatus"
	nameEntry.OnChanged = func(name string) {
		app.entity.Fields[selectedField.ID].Name = name
	}

	enumEntry := widget.NewEntry()
	enumEntry.PlaceHolder = "active, pending, done"
	enumEntry.Disable() // disable until enum type will not be selected
	enumEntry.OnChanged = func(enum string) {
		// parse enum values as string array
		app.entity.Fields[selectedField.ID].EnumValues = strings.Split(strings.ReplaceAll(enum, " ", ""), ",")
	}

	defaultEntry := widget.NewEntry()
	defaultEntry.PlaceHolder = "active"
	defaultEntry.OnChanged = func(value string) {
		app.entity.Fields[selectedField.ID].Default = value
	}

	checkEntry := widget.NewEntry()
	checkEntry.PlaceHolder = "length(field) <= 20"
	checkEntry.OnChanged = func(check string) {
		app.entity.Fields[selectedField.ID].Check = check
	}

	validationEntry := widget.NewEntry()
	validationEntry.PlaceHolder = "required"
	validationEntry.OnChanged = func(validation string) {
		app.entity.Fields[selectedField.ID].Validation = validation
	}

	referencesEntry := widget.NewEntry()
	referencesEntry.PlaceHolder = "shops (id)"
	referencesEntry.OnChanged = func(references string) {
		app.entity.Fields[selectedField.ID].References = references
	}

	fieldTypeSelect := widget.NewSelect(models.TypesString, func(selected string) {
		selectedType, ok := models.ToType(selected)
		if !ok {
			err := fmt.Errorf("type %s is not supported", selectedType)
			printer.Error(TagValidation, err)
			dialog.ShowError(err, app.window)

			return
		}

		app.entity.Fields[selectedField.ID].Type = selectedType
		if selectedType == models.TypeEnum {
			enumEntry.Enable()
		} else {
			enumEntry.Disable()
		}
	})

	nullableCheck := widget.NewCheck("Nullable", func(nullable bool) {
		app.entity.Fields[selectedField.ID].Nullable = nullable
	})

	omitemptyCheck := widget.NewCheck("Omitempty", func(omitempty bool) {
		app.entity.Fields[selectedField.ID].Omitempty = omitempty
	})

	uniqueCheck := widget.NewCheck("Unique", func(unique bool) {
		app.entity.Fields[selectedField.ID].Unique = unique
	})

	return fieldOptions{
		nameEntry:       nameEntry,
		enumEntry:       enumEntry,
		defaultEntry:    defaultEntry,
		checkEntry:      checkEntry,
		validationEntry: validationEntry,
		referencesEntry: referencesEntry,
		typeSelect:      fieldTypeSelect,
		omitemptyCheck:  omitemptyCheck,
		nullableCheck:   nullableCheck,
		uniqueCheck:     uniqueCheck,
	}
}

// generateWindow returns main Generate window.
func (a *App) generateWindow() fyne.CanvasObject {
	var (
		selectedField        = &selectedField{}
		emptyOptionContainer = container.NewCenter()
		option               = newFieldOption(a, selectedField)
	)

	return container.NewBorder(
		newEntityOptionsContainer(a),
		newGenerateButton(a),
		nil,
		nil,
		container.NewHSplit(
			newFieldsListContainer(a, selectedField, option, emptyOptionContainer),
			emptyOptionContainer,
		),
	)
}

// newFieldsListContainer returns new container that stores list of fields.
func newFieldsListContainer(
	a *App,
	selectedField *selectedField,
	option fieldOptions,
	optionsContainer fyne.CanvasObject,
) fyne.CanvasObject {
	list := widget.NewList(func() int {
		return len(a.entity.Fields)
	}, func() fyne.CanvasObject {
		return container.NewHBox(widget.NewIcon(newResource("f_icon.png")), widget.NewLabel("template"))
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*fyne.Container).Objects[1].(*widget.Label).SetText(a.entity.Fields[id].Name)
	})

	list.OnSelected = func(id widget.ListItemID) {
		selectedField.ID = id
		selectedField.selected = true

		field := a.entity.Fields[id]

		if field.Type != "" {
			option.typeSelect.Selected = field.Type.String()
		} else {
			option.typeSelect.Selected = ""
			option.typeSelect.PlaceHolder = "(Select one)"
		}

		if field.Name != defaultFieldName {
			option.nameEntry.Text = field.Name
		} else {
			option.nameEntry.Text = ""
		}

		if len(field.EnumValues) != 0 {
			option.enumEntry.Text = strings.Join(field.EnumValues, ", ")
			option.enumEntry.Enable()
		} else {
			option.enumEntry.Text = ""
			option.enumEntry.Disable()
		}

		option.defaultEntry.Text = field.Default
		option.checkEntry.Text = field.Check
		option.validationEntry.Text = field.Validation
		option.referencesEntry.Text = field.References

		option.nullableCheck.Checked = field.Nullable
		option.omitemptyCheck.Checked = field.Omitempty
		option.uniqueCheck.Checked = field.Unique

		optionsContainer.Refresh()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		selectedField.selected = false
	}

	addButton := widget.NewButton("ADD", func() {
		if len(a.entity.Fields) == 0 {
			*optionsContainer.(*fyne.Container) = newFieldOptionsContainer(option)
		}

		a.entity.Fields = append(a.entity.Fields, models.Field{Name: defaultFieldName})

		list.Refresh()
	})

	deleteButton := widget.NewButton("DELETE", func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()

		if selectedField.selected == false {
			return
		}

		a.entity.Fields = append(a.entity.Fields[:selectedField.ID], a.entity.Fields[selectedField.ID+1:]...)

		if len(a.entity.Fields) == 0 {
			*optionsContainer.(*fyne.Container) = fyne.Container{}
			optionsContainer.Refresh()
		}

		selectedField.selected = false

		list.Refresh()
	})

	listContainer := container.NewBorder(
		container.NewGridWithColumns(2, addButton, deleteButton),
		nil,
		nil,
		nil,
		list,
	)

	return listContainer
}

// newEntityOptionsContainer returns container which stores all options for entity.
func newEntityOptionsContainer(a *App) fyne.CanvasObject {
	tableEntry := widget.NewEntry()
	tableEntry.PlaceHolder = "users"
	tableEntry.OnChanged = func(result string) {
		a.entity.Table = result
	}

	packageEntry := widget.NewEntry()
	packageEntry.PlaceHolder = "users"
	packageEntry.OnChanged = func(result string) {
		a.entity.Package = result
		fmt.Println("package:", a.entity.Package)
	}

	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "user"
	nameEntry.OnChanged = func(result string) {
		a.entity.Name = result

		snake := strcase.ToSnake(result + "s")

		tableEntry.Text = snake
		a.entity.Table = snake
		tableEntry.Refresh()

		packageEntry.Text = snake
		a.entity.Package = snake
		packageEntry.Refresh()
	}

	copyrightEntry := widget.NewEntry()
	copyrightEntry.PlaceHolder = "Copyright (c)"
	copyrightEntry.OnChanged = func(copyright string) {
		a.entity.Copyright = copyright
	}

	paginationCheck := widget.NewCheck("With pagination", func(pagination bool) {
		a.entity.WithPagination = pagination
	})

	return container.New(
		layout.NewFormLayout(),
		newText("Entity", 15, color.Black, true, false),
		nameEntry,
		newText("Table", 15, color.Black, true, false),
		tableEntry,
		newText("Package", 15, color.Black, true, false),
		packageEntry,
		newText("Copyright", 15, color.Black, true, false),
		copyrightEntry,
		newText("Options", 15, color.Black, true, false),
		paginationCheck,
	)
}

// newGenerateButton returns new button to start generator.
func newGenerateButton(a *App) fyne.CanvasObject {
	newInfinite := func() *widget.ProgressBarInfinite {
		infinite := widget.NewProgressBarInfinite()
		infinite.Stop()

		return infinite
	}

	infinitesStart := func(infinites ...*widget.ProgressBarInfinite) {
		for idx := range infinites {
			infinites[idx].Start()
		}
	}

	infinitesStopAndRefresh := func(infinites ...*widget.ProgressBarInfinite) {
		for idx := range infinites {
			infinites[idx].Stop()
			infinites[idx].Refresh()
		}
	}

	infiniteOne := newInfinite()
	infiniteTwo := newInfinite()

	var generateButton *widget.Button

	generateButton = widget.NewButton("Generate", func() {
		generateButton.Disable()
		defer generateButton.Enable()

		infinitesStart(infiniteOne, infiniteTwo)
		defer infinitesStopAndRefresh(infiniteOne, infiniteTwo)

		if err := a.entity.Validate(); err != nil {
			dialog.ShowError(err, a.window)
			return
		}

		if err := a.settings.Validate(); err != nil {
			dialog.ShowError(err, a.window)
			return
		}

		a.setGenerator(generatorsInit[a.generatorType])

		if err := a.generate(); err != nil {
			printer.Error(TagGenerator, err)
			dialog.ShowError(err, a.window)
			return
		}

		dialog.ShowInformation("Success", "Generated successfully!", a.window)
	})

	generateButton.Importance = widget.HighImportance
	generateButton.Alignment = widget.ButtonAlignCenter

	return container.NewGridWithColumns(3, infiniteOne, generateButton, infiniteTwo)
}

// newFieldOptionsContainer returns new container that stores field options.
func newFieldOptionsContainer(option fieldOptions) fyne.Container {
	return *container.NewPadded(
		container.New(
			layout.NewFormLayout(),
			newText("Name", 15, color.Black, true, false),
			option.nameEntry,
			newText("Type", 15, color.Black, true, false),
			option.typeSelect,
			newText("Enum values", 15, color.Black, true, false),
			option.enumEntry,
			newText("Default", 15, color.Black, true, false),
			option.defaultEntry,
			newText("Check", 15, color.Black, true, false),
			option.checkEntry,
			newText("Validation", 15, color.Black, true, false),
			option.validationEntry,
			newText("References", 15, color.Black, true, false),
			option.referencesEntry,
			newText("Options", 15, color.Black, true, false),
			container.NewVBox(option.nullableCheck, option.omitemptyCheck, option.uniqueCheck),
		),
	)
}
