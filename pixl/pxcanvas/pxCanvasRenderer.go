package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.canvasCursor = objects
}

func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objecs := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objecs = append(objecs, &renderer.canvasBorder[i])
	}
	objecs = append(objecs, renderer.canvasImage)
	objecs = append(objecs, renderer.canvasCursor...)
	return objecs
}

func (renderer *PxCanvasRenderer) Destroy() {}

func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorders(size)
}
func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize

	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CnavasOffset.X, renderer.pxCanvas.CnavasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}
func (renderer *PxCanvasRenderer) LayoutBorders(size fyne.Size) {
	offset := renderer.pxCanvas.CnavasOffset
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}

func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pxCanvas.reloadImage = false
	}
	renderer.LayoutCanvas(renderer.pxCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}
