package main

import "fmt"

func main() {
	fmt.Println()
}

func numTimesAllBlue(light []int) int {
	res, max := 0, 0
	for i, v := range light {
		if max < v {
			max = v
		}
		if max == i+1 {
			res++
		}
	}
	return res
}
