package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swiatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)
	appLayout := container.NewBorder(nil, swiatchesContainer, nil, colorPicker, app.PixlCanvas)
	app.PixlWindow.SetContent(appLayout)
}
