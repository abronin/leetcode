package main

import "fmt"

func main() {
	input := []int{-1, -2, -3}
	fmt.Println(isMonotonic(input))
}

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	return isInc(A) || isDec(A)
}

func isInc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

func isDec(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
