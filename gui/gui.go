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
	Grid *fyne.Container
)

type Cell struct {
	rectangle *canvas.Rectangle
	text      *canvas.Text
	container *fyne.Container
}

func newCell(text string, color color.Color) (cell *Cell) {
	rect := canvas.NewRectangle(color)
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

func UpdateGrid(gridVals []int) {
	cols := 4
	value := ""
	gameGrid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < len(gridVals); i++ {

		if gridVals[i] > 0 {
			value = strconv.Itoa(gridVals[i])
		}
		cell := newCell(value, color.Gray{0x30})
		gameGrid.Add(cell.container)
		value = ""
	}
	Grid = gameGrid
}

func SetGrid() {
	cols := 4
	Grid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < cols; i++ {
		cell := newCell("", color.Gray{0x30})
		Grid.Add(cell.container)
	}

}
