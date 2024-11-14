package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType int

type PxCanvasConfig struct {
	DrawingArea  fyne.Size
	CnavasOffset fyne.Position
	PxRows       int
	PxCols       int
	PxSize       int
}

type State struct {
	BrushColor color.Color
	// BrushType      BrushType
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
