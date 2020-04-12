package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 4}, 8))
}

func twoSum(nums []int, target int) []int {
	divs := make(map[int]int)
	for i, v := range nums {
		divs[target-v] = i
	}

	for i, v := range nums {
		if val, ok := divs[v]; ok && val != i {
			return []int{i, val}
		}
	}
	return []int{-1, -1}
}
