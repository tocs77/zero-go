package pxcanvas

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize++
	case direction < 0 && pxCanvas.PxSize > 2:
		pxCanvas.PxSize--
	default:
		pxCanvas.PxSize = 10
	}
}

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	fmt.Println("Pan")
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CnavasOffset.X += xDiff
	pxCanvas.CnavasOffset.Y += yDiff
	pxCanvas.Refresh()
}
