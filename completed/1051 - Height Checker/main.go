package main

import "fmt"

func main() {
	heightChecker([]int{2, 3, 1, 5, 10, 1})
	fmt.Println()
}

func heightChecker(heights []int) int {
	sorted := sorth(cph(heights))
	return compare(heights, sorted)
}

func compare(a, b []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}
	return res
}

func cph(heights []int) []int {
	res := make([]int, len(heights))
	for i, v := range heights {
		res[i] = v
	}
	return res
}

func sorth(heights []int) []int {
	for i := 0; i < len(heights); i++ {
		index := i
		for j := i; j < len(heights); j++ {
			if heights[j] <= heights[index] {
				index = j
			}
		}
		heights[i], heights[index] = heights[index], heights[i]
	}
	return heights
}
