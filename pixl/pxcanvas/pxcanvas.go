package pxcanvas

import (
	"coursecontent/pixl/apptype"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CnavasOffset.X)
	y0 := int(pxCanvas.CnavasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + int(pxCanvas.CnavasOffset.X))
	y1 := int(pxCanvas.PxRows*pxCanvas.PxSize + int(pxCanvas.CnavasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) && pos.X < float32(bounds.Max.X) && pos.Y >= float32(bounds.Min.Y) && pos.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlancImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(appState *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       appState,
	}
	pxCanvas.PixelData = NewBlancImage(config.PxCols, config.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorders := make([]canvas.Line, 4)
	for i := 0; i < 4; i++ {
		canvasBorders[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorders[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorders,
	}

	pxCanvas.renderer = renderer
	return renderer
}

func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}
