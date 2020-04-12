package main

import (
	"fmt"
)

type Queue struct {
	values []int
}

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(input)
	fmt.Println(res)
}

func trap(height []int) int {
	start := findStart(height)
	if start == -1 {
		return 0
	}

	queue := newQueue().Push(start)
	for closer := start; closer >= 0; {
		closer = findCloser(height, start)
		queue = queue.Push(closer)
		if closer > 0 {
			queue = queue.Push(closer)
			start = closer
		}
	}

	queue, pair := queue.PullPair()
	res := 0
	for ; pair[1] >= 0; queue, pair = queue.PullPair() {
		res += calcTrap(height, pair)
	}

	res += trap(height[pair[0]+1:])

	return res
}

func calcTrap(height []int, pair []int) int {
	volume := 0
	limit := Min(height[pair[0]], height[pair[1]])

	for i := pair[0] + 1; i < pair[1]; i++ {
		volume += limit - height[i]
	}
	return volume
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func findStart(height []int) int {
	for i := 0; i < len(height)-1; i++ {
		if height[i] > 0 {
			return i
		}
	}
	return -1
}

func findCloser(height []int, start int) int {
	maxFollowerValue, maxFollowerIndex := 0, -1
	for i := start + 1; i <= len(height)-1; i++ {
		if height[i] >= height[start] {
			return i
		}
		if height[i] > maxFollowerValue {
			maxFollowerValue = height[i]
			maxFollowerIndex = i
		}
	}
	return maxFollowerIndex
}

func getContainers(height []int) Queue {
	return Queue{values: []int{1, 3, 3, 7, 10}}
}

func (queue Queue) String() string {
	return fmt.Sprintf("%v", queue.values)
}

func newQueue() Queue {
	values := make([]int, 0)
	return Queue{values: values}
}

func (queue Queue) Push(value int) Queue {
	res := queue
	res.values = append(res.values, value)
	return res
}

func (queue Queue) PullPair() (Queue, []int) {
	reduced := queue
	reduced.values = reduced.values[2:]
	pair := queue.values[:2]
	return reduced, pair
}
