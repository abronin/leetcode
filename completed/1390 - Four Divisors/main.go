package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10000
	divs := findDivs(x)
	fmt.Println(divs)
}

func sumFourDivisors(nums []int) int {
	sum := 0
	for _, x := range nums {
		divs := findDivs(x)
		if len(divs) == 4 {
			for _, div := range divs {
				sum += div
			}
		}
	}
	return sum
}

func findDivs(x int) []int {
	divs := []int{}
	for i := 1; float64(i) <= math.Sqrt(float64(x)); i++ {
		if x%i == 0 {
			divs = append(divs, i)
			if x/i != i {
				divs = append(divs, x/i)
			}
		}
	}
	return divs
}
