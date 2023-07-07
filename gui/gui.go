package gui

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"golang.org/x/image/colornames"
)

var (
	Grid     *fyne.Container
	colorMap = map[int][]uint8{
		0:    []uint8{144, 144, 144, 210},
		2:    []uint8{234, 224, 46, 210},
		4:    []uint8{234, 169, 46, 210},
		8:    []uint8{150, 234, 46, 210},
		16:   []uint8{46, 234, 193, 210},
		32:   []uint8{246, 146, 199, 210},
		64:   []uint8{245, 146, 246, 210},
		128:  []uint8{141, 125, 241, 210},
		256:  []uint8{216, 71, 14, 210},
		512:  []uint8{231, 45, 45, 210},
		1024: []uint8{70, 231, 45, 210},
		2048: []uint8{45, 231, 136, 210},
		4096: []uint8{255, 255, 0, 210},
	}
)

type Cell struct {
	rectangle *canvas.Rectangle
	text      *canvas.Text
	container *fyne.Container
}

func newCellColor(value int) (cell *Cell) {
	var cellColor []uint8
	if value >= 4096 {
		cellColor = colorMap[4096]
	} else {
		cellColor = colorMap[value]
	}
	rect := canvas.NewRectangle(color.NRGBA{R: cellColor[0], G: cellColor[1], B: cellColor[2], A: cellColor[3]})
	text := strconv.Itoa(value)
	label := canvas.NewText(text, colornames.White)
	label.TextSize = 50
	label.TextStyle = fyne.TextStyle{Bold: true}
	cell = &Cell{
		rectangle: rect,
		text:      label,
		container: container.New(layout.NewMaxLayout(), rect, label),
	}

	cell.text.Alignment = fyne.TextAlignCenter
	return cell
}

func newCell(value int, color color.Color) (cell *Cell) {

	rect := canvas.NewRectangle(color)
	text := strconv.Itoa(value)
	label := canvas.NewText(text, colornames.White)
	label.TextSize = 50
	label.TextStyle = fyne.TextStyle{Bold: true}
	cell = &Cell{
		rectangle: rect,
		text:      label,
		container: container.New(layout.NewMaxLayout(), rect, label),
	}

	cell.text.Alignment = fyne.TextAlignCenter
	return cell
}

func SetGrid() {
	cols := 4
	Grid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < cols; i++ {
		cell := newCellColor(0)
		Grid.Add(cell.container)
	}

}

func UpdateGridColor(gridVals []int) {
	cols := 4
	value := 0
	gameGrid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < len(gridVals); i++ {

		if gridVals[i] > 0 {
			value = gridVals[i]
		}
		cell := newCellColor(value)
		gameGrid.Add(cell.container)
		value = 0
	}
	Grid = gameGrid
}
