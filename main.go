package main

import (
	"2048/game"
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println("Press ESC to exit")
	game.InitBoard()
	game.ShowCLI()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		game.HandleMove(key)
		game.ShowCLI()
	}
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
