func validTicTacToe(board []string) bool {
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 'X' {
				xCount++
			} else if cell == 'O' {
				oCount++
			}
		}
	}
	if xCount != oCount && xCount != oCount+1 {
		return false
	}
	xCanWin, oCanWin := canWin(board, 'X'), canWin(board, 'O')
	if (xCanWin && oCanWin) || (xCanWin && xCount != oCount+1) || (oCanWin && xCount != oCount) {
		return false
	}
	return true
}

func canWin(board []string, player rune) bool {
	for i := 0; i < 3; i++ {
		if rune(board[i][0]) == player && rune(board[i][1]) == player && rune(board[i][2]) == player {
			return true
		}
		if rune(board[0][i]) == player && rune(board[1][i]) == player && rune(board[2][i]) == player {
			return true
		}
	}
	if rune(board[0][0]) == player && rune(board[1][1]) == player && rune(board[2][2]) == player {
		return true
	}
	if rune(board[0][2]) == player && rune(board[1][1]) == player && rune(board[2][0]) == player {
		return true
	}
	return false
}