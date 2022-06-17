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

	"crud-generator-gui/internal/models"
	"crud-generator-gui/pkg/printer"
)

// defaultFieldName returns default name for field in entity creation.
const defaultFieldName = "Field"

type (
	selectedField struct {
		id       int
		selected bool
	}

	fieldOption struct {
		nameEntry     *widget.Entry
		enumEntry     *widget.Entry
		typeSelect    *widget.Select
		nullableCheck *widget.Check
	}
)

func newFieldOption(a *App, selectedField *selectedField) fieldOption {
	fieldNameEntry := widget.NewEntry()
	fieldNameEntry.PlaceHolder = "Status"
	fieldNameEntry.OnChanged = func(s string) {
		a.entity.Fields[selectedField.id].Name = s
	}

	fieldEnumEntry := widget.NewEntry()
	fieldEnumEntry.PlaceHolder = "active, pending, done"
	fieldEnumEntry.Disable()
	fieldEnumEntry.OnChanged = func(s string) {
		a.entity.Fields[selectedField.id].EnumValues = strings.Split(strings.ReplaceAll(s, " ", ""), ",")
	}

	fieldTypeSelect := widget.NewSelect(models.TypesString, func(selected string) {
		selectedType, ok := models.ToType(selected)
		if !ok {
			err := fmt.Errorf("type %s is not supported", selectedType)
			printer.Error("TYPE VALIDATION", err)
			dialog.ShowError(err, a.window)

			return
		}

		a.entity.Fields[selectedField.id].Type = selectedType
		if selectedType == models.TypeEnum {
			fieldEnumEntry.Enable()
		} else {
			fieldEnumEntry.Disable()
		}
	})

	nullableCheck := widget.NewCheck("Nullable", func(b bool) {
		a.entity.Fields[selectedField.id].Nullable = b
	})

	return fieldOption{
		nameEntry:     fieldNameEntry,
		enumEntry:     fieldEnumEntry,
		typeSelect:    fieldTypeSelect,
		nullableCheck: nullableCheck,
	}
}

// generateWindow returns main generate window.
func (a *App) generateWindow() fyne.CanvasObject {
	var (
		selectedField        = &selectedField{}
		emptyOptionContainer = container.NewCenter()
	)

	option := newFieldOption(a, selectedField)

	return container.NewBorder(
		prepareEntityOptionsContainer(a),
		prepareGenerateButton(a),
		nil,
		nil,
		container.NewHSplit(
			prepareFieldsListContainer(a, selectedField, option, emptyOptionContainer),
			emptyOptionContainer,
		),
	)
}

func prepareFieldsListContainer(
	a *App,
	selectedField *selectedField,
	option fieldOption,
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
		selectedField.id = id
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
		} else {
			option.enumEntry.Text = ""
		}

		option.nullableCheck.Checked = field.Nullable

		optionsContainer.Refresh()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		selectedField.selected = false
	}

	addButton := widget.NewButton("ADD", func() {
		if len(a.entity.Fields) == 0 {
			*optionsContainer.(*fyne.Container) = prepareFieldOptionsContainer(option)
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

		// fmt.Println("ID:", selectedField.id)
		// fmt.Println("Selected:", selectedField.selected)

		if selectedField.selected == false {
			return
		}

		a.entity.Fields = append(a.entity.Fields[:selectedField.id], a.entity.Fields[selectedField.id+1:]...)

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

func prepareEntityOptionsContainer(a *App) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "user"
	nameEntry.OnChanged = func(result string) {
		a.entity.Name = result
	}

	tableEntry := widget.NewEntry()
	tableEntry.PlaceHolder = "users"
	tableEntry.OnChanged = func(result string) {
		a.entity.Table = result
	}

	packageEntry := widget.NewEntry()
	packageEntry.PlaceHolder = "users"
	packageEntry.OnChanged = func(result string) {
		a.entity.Package = result
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
		newText("Options", 15, color.Black, true, false),
		paginationCheck,
	)
}

func prepareGenerateButton(a *App) fyne.CanvasObject {
	button := widget.NewButton("Generate", func() {
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
			printer.Error("GENERATOR", err)
			dialog.ShowError(err, a.window)
			return
		}

		dialog.ShowInformation("Success", "Generated successfully!", a.window)
	})

	button.Alignment = widget.ButtonAlignCenter

	return container.NewGridWithColumns(3, layout.NewSpacer(), button, layout.NewSpacer())
}

func prepareFieldOptionsContainer(option fieldOption) fyne.Container {
	return *container.NewGridWithRows(
		3,
		layout.NewSpacer(),
		container.New(
			layout.NewFormLayout(),
			newText("Name", 15, color.Black, true, false),
			option.nameEntry,
			newText("Type", 15, color.Black, true, false),
			option.typeSelect,
			newText("Enum values", 15, color.Black, true, false),
			option.enumEntry,
			newText("Options", 15, color.Black, true, false),
			option.nullableCheck,
		),
		layout.NewSpacer(),
	)
}
