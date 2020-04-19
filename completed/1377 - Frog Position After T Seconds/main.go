package main

import "fmt"

func main() {
	fmt.Println()
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	levels, childs := map[int]int{1: 0}, map[int]int{}
	visited := map[int]bool{1: true}
	path := map[int][]int{}
	visit(n, 1, 0, edges, &visited, &levels, &childs, &path, []int{})

	if t > levels[target] && hasChildren(childs, target) {
		return float64(0)
	}

	if t < levels[target] {
		return float64(0)
	}

	prob := calcProb(levels[target], childs, path[target], target)
	return prob
}

func calcProb(level int, childs map[int]int, path []int, target int) float64 {
	res := float64(1)
	for i := 1; i < len(path); i++ {
		res /= float64(childs[path[i-1]])
	}
	return res
}

func visit(n, num, level int, edges [][]int, visited *map[int]bool, levels, childs *map[int]int, path *map[int][]int, pathToNum []int) {
	currentPath := make([]int, len(pathToNum)+1)
	copy(currentPath, append(pathToNum, num))
	(*path)[num] = currentPath
	(*visited)[num] = true
	(*levels)[num] = level
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if from == num && !isVisited(*visited, to) {
			_, ok := (*childs)[num]
			if !ok {
				(*childs)[num] = 0
			}
			(*childs)[num]++
			visit(n, to, level+1, edges, visited, levels, childs, path, currentPath)
		}

		if to == num && !isVisited(*visited, from) {
			_, ok := (*childs)[num]
			if !ok {
				(*childs)[num] = 0
			}
			(*childs)[num]++
			visit(n, from, level+1, edges, visited, levels, childs, path, currentPath)
		}
	}
}

func isVisited(visited map[int]bool, i int) bool {
	_, ok := visited[i]
	return ok
}

func hasChildren(childs map[int]int, num int) bool {
	_, ok := childs[num]
	return ok
}
