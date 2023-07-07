package main

import (
	"2048/game"
	"2048/gui"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("2048")
	losW := setLosW(a)
	losW.Hide()
	losW.CenterOnScreen()
	winW := setWinW(a)
	winW.Hide()
	winW.CenterOnScreen()
	game.InitBoard()
	gui.SetGrid()
	gridVals := game.Export()
	gui.UpdateGrid(gridVals)
	w.SetContent(gui.Grid)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	go func() {
		for {
			w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
				game.HandleMove(string(k.Name))
				if game.Loose {
					losW.Show()
				}
				gridVals = game.Export()
				gui.UpdateGrid(gridVals)
				w.SetContent(gui.Grid)
				if game.Win {
					winW.Show()
				}
			})
		}
	}()
	w.ShowAndRun()

}

func setLosW(a fyne.App) fyne.Window {
	losW := a.NewWindow("Defeat")
	looseTxt := canvas.NewText("You loose", color.White)
	looseTxt.TextSize = 150
	losW.SetContent(looseTxt)
	return losW
}

func setWinW(a fyne.App) fyne.Window {
	winW := a.NewWindow("Defeat")
	winTxt := canvas.NewText("You win", color.White)
	winTxt.TextSize = 150
	winW.SetContent(winTxt)
	return winW
}
