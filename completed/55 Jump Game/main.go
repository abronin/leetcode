package main

import "fmt"

func main() {
	n := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(n))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	jumpIsAvailable := true
	ns := nextStop(nums, 0, nums[0])
	start := nums[0]

	for jumpIsAvailable && ns < len(nums)-1 {
		start, ns = ns, nextStop(nums, start, ns)
		if ns == start {
			jumpIsAvailable = false
		}
	}

	return ns >= len(nums)-1
}

func nextStop(nums []int, start, end int) int {
	limit := start
	for i := start; i <= end && i <= len(nums)-1; i++ {
		if limit < i+nums[i] {
			limit = i + nums[i]
		}
	}
	return limit
}
