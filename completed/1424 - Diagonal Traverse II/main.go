package main

import "fmt"

func main() {
	fmt.Println()
}

func findDiagonalOrder(nums [][]int) []int {
	n, maxLen := len(nums), 0
	m := 0
	for _, row := range nums {
		rlen := len(row)
		m += rlen
		if maxLen < rlen {
			maxLen = rlen
		}
	}

	min, max := maxLen, n
	if n < maxLen {
		min = n
		max = maxLen
	}

	res := make([]int, m)
	stepLen := 0
	cnt := 0
	for i := 0; i < n+maxLen-1; i++ {
		if i < min {
			stepLen++
		}
		if i > max-1 {
			stepLen--
		}

		startRow := getStartRow(n, maxLen, i)
		startCol := getStartCol(n, maxLen, i)

		for j := 0; j < stepLen; j++ {
			if len(nums[startRow-j]) > startCol+j {
				el := nums[startRow-j][startCol+j]
				res[cnt] = el
				cnt++
			}
		}
	}

	return res
}

func getStartRow(n, maxLen, i int) int {
	if i < n {
		return i
	}
	return n - 1
}

func getStartCol(n, maxLen, i int) int {
	if i < n {
		return 0
	}
	return i + 1 - n
}
