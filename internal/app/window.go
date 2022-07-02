package app

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// mainWindow is the main application gui window.
func (a *App) mainWindow() fyne.Window {
	fyneApplication := fyneApp.New()
	// a.Settings().SetTheme(theme.DarkTheme())

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Generate", newResource("go_icon.png"), a.generateWindow()),
		container.NewTabItemWithIcon("Settings", newResource("settings_icon.png"), a.settingsWindow()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	a.window = fyneApplication.NewWindow(a.meta.AppName)
	a.window.SetContent(tabs)
	a.window.CenterOnScreen()
	a.window.Resize(fyne.NewSize(900, 600))
	a.window.SetIcon(newResource("go_icon.png"))

	return a.window
}
