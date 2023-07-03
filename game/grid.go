package game

import (
	"math/rand"
)

type grid struct {
	tiles [][]tile
	rows  int
	cols  int
}

type tile struct {
	value int
	location
}

type location struct {
	col int
	row int
}

func StartGame() {
	grid := startGrid(4, 4)
	grid.setupGame()
}

func startGrid(rows, cols int) grid {
	var grid grid
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			grid.tiles[i][j] = tile{
				value: 0,
				location: location{
					col: i,
					row: j,
				},
			}
		}
	}
	grid.rows = rows
	grid.cols = cols
	return grid
}

func (g *grid) setupGame() {
	g.placeInFreeTile()
	g.placeInFreeTile()
}

func (g *grid) placeInFreeTile() {
	freeTiles := g.getFreeTiles()
	toPlace := freeTiles[rand.Intn(len(freeTiles))]
	toPlace.value = 2
	g.tiles[toPlace.col][toPlace.row] = toPlace
}

func (g *grid) getFreeTiles() []tile {
	var freeTiles []tile
	for i := 0; i < len(g.tiles); i++ {
		for j := 0; j < len(g.tiles[i]); j++ {
			if g.tiles[i][j].value == 0 {
				freeTiles = append(freeTiles, g.tiles[i][j])
			}
		}
	}
	return freeTiles
}

func (t *tile) absorb(another tile) bool {
	if t.value == another.value {
		t.value += another.value
		return true
	}
	return false
}

func (g *grid) moveLeft() {
	g.smooshLeft()
	for j := 0; j < g.rows-1; j++ {
		for i := 0; i < g.cols; i++ {
			ok := g.tiles[i][j].absorb(g.tiles[i+1][j])
			if ok {
				g.smooshLeft()
				i--
			}
		}
	}
}

func (g *grid) smooshLeft() {
	for j := 0; j < g.rows; j++ {
		count := 0
		for i := 0; i < g.cols; i++ {
			currTile := g.tiles[i][j]
			if currTile.value > 0 {
				g.tiles[count][j] = currTile
				currTile.value = 0
				count++
			}
		}
	}
}
