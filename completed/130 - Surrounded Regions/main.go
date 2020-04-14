package main

import "fmt"

func main() {
	input := [][]byte{
		{'X', 'X', 'O'},
		{'X', 'O', 'X'},
		{'O', 'X', 'X'},
	}
	solve(input)
	fmt.Printf("%c\n", input)
}

func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])

	for i := 0; i < m; i++ {
		markD(board, 0, i, n, m)
		markD(board, n-1, i, n, m)
	}

	for i := 0; i < n; i++ {
		markD(board, i, 0, n, m)
		markD(board, i, m-1, n, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != 'D' {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}
}

func markD(board [][]byte, x, y, n, m int) {
	if board[x][y] != 'O' {
		return
	}

	board[x][y] = 'D'

	if x != 0 {
		markD(board, x-1, y, n, m)
	}

	if x != n-1 {
		markD(board, x+1, y, n, m)
	}

	if y != 0 {
		markD(board, x, y-1, n, m)
	}

	if y != m-1 {
		markD(board, x, y+1, n, m)
	}
}
