package main

import "fmt"

func main() {
	fmt.Println("Run tests!")
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	distance := 0
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			v := arr1[i] - arr2[j]
			if v < 0 {
				v = -v
			}
			if v <= d {
				flag = false
			}
		}
		if flag {
			distance++
		}
	}
	return distance
}
