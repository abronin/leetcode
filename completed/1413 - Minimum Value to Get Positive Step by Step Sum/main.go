package main

import "fmt"

func main() {
	input := []int{-3, 2, -3, 4, 2}
	fmt.Println(minStartValue(input))
}

func minStartValue(nums []int) int {
	min, sum := -100*-100-1, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if min > sum {
			min = sum
		}
	}
	if min >= 0 {
		return 1
	}
	return -min + 1
}
