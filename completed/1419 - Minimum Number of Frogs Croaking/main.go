package main

import (
	"fmt"
	// "github.com/pkg/profile"
)

func main() {
	// defer profile.Start().Stop()
	input := "crcoakroak"
	fmt.Println(minNumberOfFrogs(input))
}

func minNumberOfFrogs(croakOfFrogs string) int {
	sequence := make([]int, 5)
	max := 0
	for _, v := range croakOfFrogs {
		frogs := count(&sequence, v)
		if frogs < 0 {
			return -1
		}
		if max < frogs {
			max = frogs
		}
	}

	if !full(&sequence) {
		return -1
	}

	return max
}

func count(seq *[]int, c rune) int {
	switch c {
	case 'c':
		i := 0
		(*seq)[i]++
	case 'r':
		i := 1
		(*seq)[i]++
		if (*seq)[i-1] < (*seq)[i] {
			return -1
		}
	case 'o':
		i := 2
		(*seq)[i]++
		if (*seq)[i-1] < (*seq)[i] {
			return -1
		}
	case 'a':
		i := 3
		(*seq)[i]++
		if (*seq)[i-1] < (*seq)[i] {
			return -1
		}
	case 'k':
		i := 4
		(*seq)[i]++
		if (*seq)[i-1] < (*seq)[i] {
			return -1
		}
	}
	return (*seq)[0] - (*seq)[4]
}

func full(seq *[]int) bool {
	return (*seq)[4] == (*seq)[0]
}
