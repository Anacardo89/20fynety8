package game

func (t *tile) absorb(another tile) bool {
	if t.value == another.value {
		t.value += another.value
		return true
	}
	return false
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

func (g *grid) moveLeft() {
	g.smooshLeft()
	for j := 0; j < g.rows; j++ {
		for i := 0; i < g.cols-1; i++ {
			ok := g.tiles[i][j].absorb(g.tiles[i+1][j])
			if ok {
				g.smooshLeft()
				i--
			}
		}
	}
}

func (g *grid) smooshRight() {
	for j := g.rows - 1; j >= 0; j-- {
		count := g.rows - 1
		for i := g.cols - 1; i >= 0; i-- {
			currTile := g.tiles[i][j]
			if currTile.value > 0 {
				g.tiles[count][j] = currTile
				currTile.value = 0
				count--
			}
		}
	}
}

func (g *grid) moveRight() {
	g.smooshRight()
	for j := g.rows - 1; j >= 0; j-- {
		for i := g.cols - 1; i > 0; i-- {
			ok := g.tiles[i][j].absorb(g.tiles[i-1][j])
			if ok {
				g.smooshRight()
				i++
			}
		}
	}
}

func (g *grid) smooshUp() {
	for i := 0; i < g.cols; i++ {
		count := 0
		for j := 0; j < g.rows; j++ {
			currTile := g.tiles[i][j]
			if currTile.value > 0 {
				g.tiles[i][count] = currTile
				currTile.value = 0
				count++
			}
		}
	}
}

func (g *grid) moveUp() {
	g.smooshUp()
	for i := 0; i < g.cols; i++ {
		for j := 0; j < g.rows-1; j++ {
			ok := g.tiles[i][j].absorb(g.tiles[i][j+1])
			if ok {
				g.smooshUp()
				i--
			}
		}
	}
}

func (g *grid) smooshDown() {
	for i := g.cols - 1; i >= 0; i-- {
		count := g.cols - 1
		for j := g.rows - 1; j >= 0; j-- {
			currTile := g.tiles[i][j]
			if currTile.value > 0 {
				g.tiles[i][count] = currTile
				currTile.value = 0
				count--
			}
		}
	}
}

func (g *grid) moveDown() {
	g.smooshDown()
	for i := g.cols - 1; i >= 0; i-- {
		for j := g.rows - 1; j > 0; j-- {
			ok := g.tiles[i][j].absorb(g.tiles[i][j-1])
			if ok {
				g.smooshDown()
				i++
			}
		}
	}
}
