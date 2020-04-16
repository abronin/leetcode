package main

import "fmt"

func main() {
	fmt.Println("Run tests!")
}

func getKth(lo int, hi int, k int) int {
	values, powers := make([]int, hi-lo+1), make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		values[i-lo] = i
		powers[i-lo] = power(i)
	}

	st(values, powers)
	return values[k-1]
}

func st(values, powers []int) {
	for k := 0; k < len(values); k++ {
		mini, minv := len(values), 1<<32
		for i := k; i < len(values); i++ {
			if minv >= powers[i] {
				if mini == len(values) {
					mini = i
				}
				if minv > powers[i] || values[mini] > values[i] {
					mini = i
				}
				minv = powers[i]
			}
		}
		values[k], values[mini] = values[mini], values[k]
		powers[k], powers[mini] = powers[mini], powers[k]
	}
}

func step(x int) int {
	if x%2 == 0 {
		return int(x / 2)
	}
	return x*3 + 1
}

func power(x int) int {
	p := 0
	for x > 1 {
		x = step(x)
		p++
	}
	return p
}
