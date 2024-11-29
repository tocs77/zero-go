package ui

import (
	"coursecontent/pixl/util"
	"errors"
	"image/png"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func buildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{
			widthFormEntry,
			heightFormEntry,
		}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if !ok {
				return
			}

			pixelWidth := 0
			pixelHeight := 0
			if widthEntry.Validate() != nil {
				dialog.ShowError(errors.New("invalid width"), app.PixlWindow)
			} else {
				pixelWidth, _ = strconv.Atoi(widthEntry.Text)
			}
			if heightEntry.Validate() != nil {
				dialog.ShowError(errors.New("invalid height"), app.PixlWindow)
			} else {
				pixelHeight, _ = strconv.Atoi(heightEntry.Text)
			}
			app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)

		}, app.PixlWindow)
	})
}

func buildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu(
		"File",
		buildNewMenu(app),
		buildSaveMenu(app),
		buildSaveAsMenu(app),
		buildOpenMenu(app),
	)
}

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, err error) {
		if uri == nil {
			return
		}
		if err != nil {
			dialog.ShowError(err, app.PixlWindow)
		}
		e := png.Encode(uri, app.PixlCanvas.PixelData)
		if e != nil {
			dialog.ShowError(e, app.PixlWindow)
		}
		app.State.SetFilePath(uri.URI().Path())
	}, app.PixlWindow)
}

func buildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		saveFileDialog(app)
	})
}

func buildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PixlWindow)
				}
			}

			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
			err = png.Encode(fh, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
		}
	})
}

func buildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if uri == nil {
				return
			}
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
			img, err := png.Decode(uri)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
			app.PixlCanvas.LoadImage(img)
			app.State.SetFilePath(uri.URI().Path())
			imgColors := util.GetImageColors(img)
			i := 0
			for c := range imgColors {
				if i == len(app.Swatches) {
					break
				}
				app.Swatches[i].SetColor(c)
				i++
			}
		}, app.PixlWindow)
	})
}

func SetupMenus(app *AppInit) {
	menus := buildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWindow.SetMainMenu(mainMenu)
}

func sizeValidator(s string) error {
	width, err := strconv.Atoi(s)
	if err != nil {
		return errors.New("width must be an integer")
	}
	if width < 0 {
		return errors.New("width must be greater than 0")
	}
	return nil
}
