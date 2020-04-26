package main

import "fmt"

func main() {
	fmt.Println()
}

func maxScore(s string) int {
	n, count := len(s), 0

	if n == 1 {
		return 1
	}

	zeros := make([]int, n)
	for i, c := range s {
		if c == '0' {
			count++
		}
		zeros[i] = count
	}

	max := 0
	o := zeros[n-1]
	for i := 0; i < len(s)-1; i++ {
		z := zeros[i]
		ones := n - (i + 1) - (o - z)
		if max < z+ones {
			max = z + ones
		}
		fmt.Println("Zeros:", z, "Ones:", ones, "Max:", max)
	}
	return max
}
