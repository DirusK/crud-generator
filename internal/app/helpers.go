package app

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// newText returns new text layout.
func newText(text string, size float32, color color.Color, bold, italic bool) fyne.CanvasObject {
	projectText := canvas.NewText(text, color)
	projectText.TextSize = size
	projectText.TextStyle = fyne.TextStyle{Bold: bold, Italic: italic}

	return projectText
}
