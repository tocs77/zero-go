package ui

import (
	"coursecontent/pixl/apptype"
	"coursecontent/pixl/pxcanvas"
	"coursecontent/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
	PixlCanvas *pxcanvas.PxCanvas
}
