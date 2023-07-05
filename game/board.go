package game

import (
	"fmt"
	"math/rand"
)

var (
	board [][]int
	rows  = 4
	cols  = 4
)

type point struct {
	x int
	y int
}

func InitBoard() {
	for i := 0; i < rows; i++ {
		var col []int
		for j := 0; j < cols; j++ {
			col = append(col, 0)
		}
		board = append(board, col)
	}
	placeInFreeTile()
	placeInFreeTile()
}

func getFreeTiles() []point {
	var freeTiles []point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 0 {
				p1 := point{
					x: i,
					y: j,
				}
				freeTiles = append(freeTiles, p1)

			}
		}

	}
	return freeTiles
}

func placeInFreeTile() {
	freeTiles := getFreeTiles()
	toPlace := freeTiles[rand.Intn(len(freeTiles))]
	board[toPlace.x][toPlace.y] = 2
}

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
