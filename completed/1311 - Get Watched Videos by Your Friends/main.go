package main

import "fmt"

func main() {
	fmt.Println()
}

func findLevel(friends [][]int, id, level int) []int {
	visited := map[int]int{id: 0}
	current := map[int]bool{}
	for i := 0; i < level; i++ {
		current = map[int]bool{}
		for user, ufriends := range friends {
			if l, ok := visited[user]; ok && l == i {
				for _, friend := range ufriends {
					if _, fok := visited[friend]; !fok {
						current[friend] = true
						visited[friend] = i + 1
					}
				}
			}
		}
	}

	res := []int{}
	for i := range current {
		res = append(res, i)
	}
	return res
}

func getFilms(leveled []int, watchedVideos [][]string) []string {
	films := []string{}
	for _, u := range leveled {
		for _, film := range watchedVideos[u] {
			films = append(films, film)
		}
	}
	return films
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	leveled := findLevel(friends, id, level)
	levels := map[int]int{id: 0}
	for i := 0; i < level; i++ {
		for user, frofu := range friends {
			if ulevel, ok := levels[user]; ok {
				for _, frnd := range frofu {
					if _, okok := levels[frnd]; !okok {
						levels[frnd] = ulevel + 1
					}
				}
			}
		}
	}

	films := getFilms(leveled, watchedVideos)

	freq := map[string]int{}
	for _, film := range films {
		if _, ok := freq[film]; !ok {
			freq[film] = 0
		}
		freq[film]++
	}

	res := []string{}
	for len(freq) > 0 {
		min := 100001
		mins := []string{}
		for film, f := range freq {
			if min >= f {
				if min > f {
					mins = []string{film}
				} else if min == f {
					mins = append(mins, film)
				}
				min = f
			}
		}

		current := minimal(mins)
		delete(freq, current)
		res = append(res, current)
	}

	return res
}

func minimal(mins []string) string {
	m := mins[0]
	for _, v := range mins {
		if v < m {
			m = v
		}
	}
	return m
}
