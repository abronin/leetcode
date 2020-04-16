package main

import "fmt"

func main() {
	fmt.Println("Run tests!")
}

func findLucky(arr []int) int {
	fr := map[int]int{}
	for _, x := range arr {
		if _, ok := fr[x]; !ok {
			fr[x] = 1
		} else {
			fr[x]++
		}
	}

	res := -1
	for i, v := range fr {
		if i == v && res < i {
			res = i
		}
	}

	return res
}
