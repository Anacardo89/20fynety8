package main

import (
	"2048/game"
	"bufio"
	"fmt"
	"os"
)

func main() {
	game.InitBoard()
	game.ShowCLI()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Move: ")
		move, _ := reader.ReadString('\n')
		m := string(move[0])
		game.HandleMove(m)
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
