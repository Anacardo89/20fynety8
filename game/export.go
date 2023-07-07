package game

import "fmt"

func ShowCLI() {
	for i := 0; i < len(board); i++ {
		fmt.Print("|")
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
			fmt.Print("|")
		}
		fmt.Println()
	}
	fmt.Println()
}

func Export() []int {
	var grid []int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid = append(grid, board[i][j])
			checkWin(i, j)
		}
	}
	return grid
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
