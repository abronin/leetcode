package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(maxDiff(1234))
	val := chSym("12343", 3, 9)
	fmt.Println(val)
}

func maxDiff(num int) int {
	s := strconv.Itoa(num)
	notZeroNum := -1
	notNineNum := -1
	for i := 0; i < len(s); i++ {
		posnum := getpos(s, i)
		if posnum > 1 && notZeroNum < 0 {
			notZeroNum = posnum
		}
		if posnum < 9 && notNineNum < 0 {
			notNineNum = posnum
		}
	}

	max := chSym(s, notNineNum, 9)
	min := 0
	if getpos(s, 0) == 1 {
		min = chSym(s, notZeroNum, 0)
	} else {
		min = chSym(s, getpos(s, 0), 1)
	}

	return max - min
}

func getpos(s string, pos int) int {
	i, _ := strconv.Atoi(s[pos : pos+1])
	return i
}

func chSym(s string, from, to int) int {
	fromR := rune('0' + from)
	toR := rune('0' + to)
	res := ""
	for _, c := range s {
		if c == fromR {
			res += string(toR)
		} else {
			res += string(c)
		}
	}
	i, _ := strconv.Atoi(res)
	return i
}
