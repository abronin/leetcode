package main

import "fmt"

func main() {
	// input := []int{1, 2, 0}
	// input := []int{2, 1, 0}
	// input := []int{7, 8, 9, 11, 12}
	input := []int{3, 4, -1, 4}
	res := firstMissingPositive(input)
	fmt.Println(res)
}

func firstMissingPositive(nums []int) int {
	length := len(nums)
	msdRadixSort(&nums)
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return length + 1
}

func msdRadixSort(numsP *[]int) {
	length := len(*numsP)
	for i := 0; i < length; i++ {
		if (*numsP)[i] > length {
			(*numsP)[i] = 0
		}

		if (*numsP)[i] > 0 {
			if (*numsP)[i] != i+1 {
				interim(numsP, i, (*numsP)[i]-1)
				i--
				continue
			}
		} else {
			(*numsP)[i] = 0
		}
	}
}

func interim(numsP *[]int, a int, b int) {
	if (*numsP)[a] == (*numsP)[b] {
		(*numsP)[b] = 0
	}
	(*numsP)[a], (*numsP)[b] = (*numsP)[b], (*numsP)[a]
}
