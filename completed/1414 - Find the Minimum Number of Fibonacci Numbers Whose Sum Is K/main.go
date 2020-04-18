package main

import "fmt"

func main() {
	fmt.Println()
}

func findMinFibonacciNumbers(k int) int {
	fib, current := []int{1, 1}, 1

	for fib[current] <= k {
		fib = append(fib, fib[current]+fib[current-1])
		current++
	}

	res := 0
	for i := current; k > 0; i-- {
		for k >= fib[i] {
			res++
			k -= fib[i]
		}
	}

	return res
}
