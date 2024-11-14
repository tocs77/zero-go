package swatch

import (
	"coursecontent/pixl/apptype"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, c color.Color, index int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{Color: c, SwatchIndex: index, clickHandler: clickHandler, Selected: false}
	s.ExtendBaseWidget(s)
	return s
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{square: square, objects: objects, parent: swatch}
}
