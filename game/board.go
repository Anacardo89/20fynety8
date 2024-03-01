package game

import (
	"math/rand"
)

var (
	board    [][]int
	tmpBoard [][]int
	rows     = 4
	cols     = 4
	Win      = false
	Loose    = false
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

func InitBoardLoss() {
	count := 1
	for i := 0; i < rows; i++ {
		var col []int
		for j := 0; j < cols; j++ {
			count++
			col = append(col, count)
		}
		board = append(board, col)
	}
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
	if len(freeTiles) == 0 {
		Loose = true
		return
	}
	toPlace := freeTiles[rand.Intn(len(freeTiles))]
	board[toPlace.x][toPlace.y] = 2
}

func checkWin(i, j int) {
	if board[i][j] == 2048 {
		Win = true
		return
	}

}
