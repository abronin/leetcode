package main

import "fmt"

func main() {
	fmt.Println()
}

func xorQueries(arr []int, queries [][]int) []int {
	n := len(queries)
	res := make([]int, n)
	for i, query := range queries {
		res[i] = arr[query[0]]
		for j := query[0] + 1; j <= query[1]; j++ {
			res[i] = res[i] ^ arr[j]
		}
	}
	return res
}
