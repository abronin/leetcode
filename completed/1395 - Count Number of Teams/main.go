package main

import "fmt"

func main() {
	fmt.Println("Run tests!")
}

func numTeams(rating []int) int {
	inc, dec := 0, 0

	for i := 0; i < len(rating); i++ {
		for j := i; j < len(rating); j++ {
			for k := j; k < len(rating); k++ {
				if rating[i] < rating[j] && rating[j] < rating[k] {
					inc++
				}
				if rating[i] > rating[j] && rating[j] > rating[k] {
					dec++
				}
			}
		}
	}
	return inc + dec
}
