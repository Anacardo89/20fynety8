package main

import (
	"2048/game"
	"2048/gui"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("2048")
	game.InitBoard()
	gui.SetGrid()
	gridVals := game.Export()
	gui.UpdateGrid(gridVals)
	w.SetContent(gui.Grid)
	go func() {
		for {
			w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
				game.HandleMove(string(k.Name))
				gridVals = game.Export()
				gui.UpdateGrid(gridVals)
				w.SetContent(gui.Grid)
				for i := 0; i < len(gridVals); i++ {
					if gridVals[i] == 2048 {
						winTxt := canvas.NewText("You win", color.White)
						winTxt.TextSize = 200
						w.SetContent(winTxt)

					}
				}
			})
		}
	}()
	w.ShowAndRun()

}

func ShowIndex(rows, cols int) {
	for i := 0; i < rows; i++ {
		fmt.Print("|")
		for j := 0; j < cols; j++ {
			fmt.Print(i, j)
			fmt.Print("|")
		}
		fmt.Println()
	}
	fmt.Println()
}
