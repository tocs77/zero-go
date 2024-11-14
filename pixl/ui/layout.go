package ui

func Setup(app *AppInit) {
	swiatchesContainer := BuildSwatches(app)
	app.PixlWindow.SetContent((swiatchesContainer))
}
