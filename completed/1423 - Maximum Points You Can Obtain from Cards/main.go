package main

import "fmt"

func main() {
	fmt.Println()
}

func maxScore(cardPoints []int, k int) int {
	head, tail := make([]int, k), make([]int, k)
	headSum, tailSum := 0, 0
	n := len(cardPoints)

	if k == 1 {
		if cardPoints[0] > cardPoints[n-1] {
			return cardPoints[0]
		}
		return cardPoints[n-1]
	}

	for i := 0; i < k; i++ {
		headSum += cardPoints[i]
		tailSum += cardPoints[n-i-1]
		head[i] = headSum
		tail[i] = tailSum
	}

	var max int
	if headSum > tailSum {
		max = headSum
	} else {
		max = tailSum
	}

	for i := 0; i < k-1; i++ {
		sum := head[i] + tail[k-i-2]
		if max < sum {
			max = sum
		}
	}

	return max
}
