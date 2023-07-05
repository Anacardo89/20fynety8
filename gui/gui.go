package gui

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"golang.org/x/image/colornames"
)

var (
	App    fyne.App
	Window fyne.Window
	lout   *fyne.Container
	grid   *fyne.Container
	text   *canvas.Text
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

func SetupGUI() {
	App = app.New()
	Window = App.NewWindow("2048")
	Window.Resize(fyne.NewSize(600, 600))
	setLayout(startGrid())
	Window.SetContent(lout)
}

func Update(gridVals []int) {
	grid = updateGrid(gridVals)
	setLayout(grid)
	Window.SetContent(lout)
}

func setLayout(grid *fyne.Container) {
	text = canvas.NewText("Use the Arrow keys to move", color.White)
	lout = container.New(layout.NewVBoxLayout(), grid, text)
}

func startGrid() *fyne.Container {
	cols := 4
	gameGrid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < cols*cols; i++ {
		cell := newCell(strconv.Itoa(0), color.Gray{0x30})
		gameGrid.Add(cell.container)
	}
	return gameGrid
}

func updateGrid(gridVals []int) *fyne.Container {
	cols := 4
	gameGrid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < len(gridVals); i++ {
		cell := newCell(strconv.Itoa(gridVals[i]), color.Gray{0x30})
		gameGrid.Add(cell.container)
	}
	return gameGrid
}
