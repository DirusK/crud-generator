package app

import (
	"embed"
	"image/color"
	"io/fs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"crud-generator/pkg/printer"
)

//go:embed assets
var assets embed.FS

// newText returns new text layout.
func newText(text string, size float32, color color.Color, bold, italic bool) fyne.CanvasObject {
	projectText := canvas.NewText(text, color)
	projectText.TextSize = size
	projectText.TextStyle = fyne.TextStyle{Bold: bold, Italic: italic}

	return projectText
}

// newResource returns new static resource
func newResource(name string) *fyne.StaticResource {
	file, err := fs.ReadFile(assets, "assets/"+name)
	if err != nil {
		printer.Fatal(TagGUI, err)
	}

	return fyne.NewStaticResource(name, file)
}

// split function by `/` and '\' symbols.
func split(r rune) bool {
	return r == '/' || r == '\\'
}
