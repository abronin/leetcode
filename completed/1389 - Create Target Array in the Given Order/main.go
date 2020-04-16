package main

import "fmt"

func main() {
	fmt.Println()
}

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		for j := len(index) - 2; j >= index[i]; j-- {
			res[j+1] = res[j]
		}
		res[index[i]] = v
	}
	return res
}
