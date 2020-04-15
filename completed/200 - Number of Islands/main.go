package main

import "fmt"

func main() {
	input := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	fmt.Println(numIslands(input))
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if delIsland(grid, n, m, i, j) {
				islands++
			}
		}
	}
	return islands
}

func delIsland(grid [][]byte, n, m, x, y int) bool {
	if grid[x][y] == byte("1"[0]) {
		grid[x][y] = byte("0"[0])
		if x > 0 {
			delIsland(grid, n, m, x-1, y)
		}
		if x < n-1 {
			delIsland(grid, n, m, x+1, y)
		}
		if y > 0 {
			delIsland(grid, n, m, x, y-1)
		}
		if y < m-1 {
			delIsland(grid, n, m, x, y+1)
		}
		return true
	}
	return false
}

func intsToBytes(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = byte("1"[0])
			} else {
				grid[i][j] = byte("0"[0])
			}
		}
	}
}
