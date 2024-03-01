package ui

import (
	"20fynety8/game"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"golang.org/x/image/colornames"
)

func RunApp() {
	a := app.New()
	w := a.NewWindow("20fynety8")
	lw := setLooseWindow(a)
	ww := setWinWindow(a)
	game.InitBoard()
	SetGrid()
	gridVals := game.Export()
	UpdateGridColor(gridVals)
	w.SetContent(Grid)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	go func() {
		for {
			w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
				game.HandleMove(string(k.Name))
				if game.Loose {
					lw.Show()
				}
				gridVals = game.Export()
				UpdateGridColor(gridVals)
				w.SetContent(Grid)
				if game.Win {
					ww.Show()
				}
			})
		}
	}()
	w.ShowAndRun()
}

func setLooseWindow(a fyne.App) fyne.Window {
	losW := a.NewWindow("Defeat")
	looseTxt := canvas.NewText("You loose", color.White)
	looseTxt.TextSize = 150
	losW.SetContent(looseTxt)
	losW.Hide()
	losW.CenterOnScreen()
	return losW
}

func setWinWindow(a fyne.App) fyne.Window {
	winW := a.NewWindow("Victory")
	winTxt := canvas.NewText("You win", color.White)
	winTxt.TextSize = 150
	winW.SetContent(winTxt)
	winW.Hide()
	winW.CenterOnScreen()
	return winW
}

var (
	Grid     *fyne.Container
	colorMap = map[int][]uint8{
		0:    {144, 144, 144, 210},
		2:    {234, 224, 46, 210},
		4:    {234, 169, 46, 210},
		8:    {150, 234, 46, 210},
		16:   {46, 234, 193, 210},
		32:   {246, 146, 199, 210},
		64:   {245, 146, 246, 210},
		128:  {141, 125, 241, 210},
		256:  {216, 71, 14, 210},
		512:  {231, 45, 45, 210},
		1024: {70, 231, 45, 210},
		2048: {45, 231, 136, 210},
		4096: {255, 255, 0, 210},
	}
)

type Cell struct {
	rectangle *canvas.Rectangle
	text      *canvas.Text
	container *fyne.Container
}

func SetGrid() {
	cols := 4
	Grid := container.New(layout.NewGridLayout(cols))
	for i := 0; i < cols; i++ {
		cell := newCell(0)
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
		cell := newCell(value)
		gameGrid.Add(cell.container)
		value = 0
	}
	Grid = gameGrid
}

func newCell(value int) (cell *Cell) {
	var cellColor []uint8
	if value >= 4096 {
		cellColor = colorMap[4096]
	} else {
		cellColor = colorMap[value]
	}
	rect := canvas.NewRectangle(color.NRGBA{R: cellColor[0], G: cellColor[1], B: cellColor[2], A: cellColor[3]})
	rect.CornerRadius = 15
	text := strconv.Itoa(value)
	label := canvas.NewText(text, colornames.White)
	label.TextSize = 50
	label.TextStyle = fyne.TextStyle{Bold: true}
	cell = &Cell{
		rectangle: rect,
		text:      label,
		container: container.New(layout.NewStackLayout(), rect, label),
	}
	cell.text.Alignment = fyne.TextAlignCenter
	return cell
}
