package main

import "fmt"

func main() {
	input := []int{12, 1, 12}
	res := kidsWithCandies(input, 10)
	fmt.Println(res)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	if n == 0 {
		return []bool{}
	}
	if n == 1 {
		return []bool{true}
	}

	max := 0
	for _, i := range candies {
		if max < i {
			max = i
		}
	}

	res := make([]bool, n)
	for i, c := range candies {
		res[i] = c+extraCandies >= max
	}
	return res
}
