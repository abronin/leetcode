package main

import "fmt"

func main() {
	fmt.Println(1 << 0)
}

func hammingDistance(x int, y int) int {
	distance := 0
	for i := 30; i >= 0; i-- {
		distanceFlag := 0

		if x >= 1<<i {
			x = x - 1<<i
			distanceFlag++
		}

		if y >= 1<<i {
			y = y - 1<<i
			distanceFlag++
		}

		if distanceFlag == 1 {
			distance++
		}
	}
	return distance
}
