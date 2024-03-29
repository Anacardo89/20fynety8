package game

func HandleMove(input string) {
	tmpBoard = make([][]int, len(board))
	for i := range board {
		tmpBoard[i] = make([]int, len(board[i]))
		copy(tmpBoard[i], board[i])
	}
	switch input {
	case "Left":
		moveLeft()
	case "Right":
		moveRight()
	case "Up":
		moveUp()
	case "Down":
		moveDown()
	default:
	}
	if !equalBoards() {
		placeInFreeTile()
	} else {
		freeTiles := getFreeTiles()
		if len(freeTiles) == 0 {
			Loose = true
			return
		}
	}
}

func absorb(i, j, x, y int) bool {
	if board[i][j] == board[x][y] && board[i][j] > 0 {
		board[i][j] += board[x][y]
		board[x][y] = 0
		return true
	}
	return false
}

func equalBoards() bool {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != tmpBoard[i][j] {
				return false
			}
		}
	}
	return true
}

func moveLeft() {
	for i := 0; i < rows; i++ {
		smooshLeft(i)
		for j := 0; j < cols; j++ {
			if j < cols-1 {
				ok := absorb(i, j, i, j+1)
				if ok {
					smooshLeft(i)
				}
			}
		}
	}
}

func smooshLeft(i int) {
	count := 0
	for j := 0; j < cols; j++ {
		curr := board[i][j]
		if curr > 0 {
			board[i][count] = curr
			if j != count {
				board[i][j] = 0
			}
			count++
		}
	}
}

func moveRight() {
	for i := 0; i < rows; i++ {
		smooshRight(i)
		for j := cols - 1; j >= 0; j-- {
			if j > 0 {
				ok := absorb(i, j, i, j-1)
				if ok {
					smooshRight(i)
				}
			}
		}
	}
}

func smooshRight(i int) {
	count := cols - 1
	for j := cols - 1; j >= 0; j-- {
		curr := board[i][j]
		if curr > 0 {
			board[i][count] = curr
			if j != count {
				board[i][j] = 0
			}
			count--
		}
	}
}

func moveUp() {
	for i := 0; i < cols; i++ {
		smooshUp(i)
		for j := 0; j < rows; j++ {
			if j < rows-1 {
				ok := absorb(j, i, j+1, i)
				if ok {
					smooshUp(i)
				}
			}
		}
	}
}

func smooshUp(i int) {
	count := 0
	for j := 0; j < rows; j++ {
		curr := board[j][i]
		if curr > 0 {
			board[count][i] = curr
			if j != count {
				board[j][i] = 0
			}
			count++
		}
	}
}

func moveDown() {
	for i := cols - 1; i >= 0; i-- {
		smooshDown(i)
		for j := rows - 1; j >= 0; j-- {
			if j > 0 {
				ok := absorb(j, i, j-1, i)
				if ok {
					smooshDown(i)
				}
			}
		}
	}
}

func smooshDown(i int) {
	count := rows - 1
	for j := rows - 1; j >= 0; j-- {
		curr := board[j][i]
		if curr > 0 {
			board[count][i] = curr
			if j != count {
				board[j][i] = 0
			}
			count--
		}
	}
}
