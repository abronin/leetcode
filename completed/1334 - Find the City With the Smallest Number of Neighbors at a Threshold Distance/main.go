package main

import "fmt"

func main() {
	fmt.Println()
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	count := map[int]int{}
	for start := 0; start < n; start++ {
		distances := countDistances(n, edges, distanceThreshold, start)
		count[start] = len(distances) - 1
	}

	return maxIDOfMinCount(n, count)
}

func countDistances(n int, edges [][]int, distanceThreshold, start int) map[int]int {
	distances := map[int]int{start: 0}
	for i := 0; i < n; i++ {
		for _, edge := range edges {
			from, to, distance := edge[0], edge[1], edge[2]
			distances = visitEdge(distances, from, to, distance, distanceThreshold)
			distances = visitEdge(distances, to, from, distance, distanceThreshold)
		}
	}
	return distances
}

func visitEdge(distances map[int]int, from, to, distance, distanceThreshold int) map[int]int {
	if dFrom, visFrom := distances[from]; visFrom {
		if dTo, visTo := distances[to]; !visTo && dFrom+distance <= distanceThreshold {
			distances[to] = dFrom + distance
		} else if visTo && dFrom+distance < dTo {
			distances[to] = dFrom + distance
		}
	}
	return distances
}

func maxIDOfMinCount(n int, count map[int]int) int {
	min := 0
	for index, neighbors := range count {
		if count[min] > neighbors || (count[min] == neighbors && min < index) {
			min = index
		}
	}
	return min
}
