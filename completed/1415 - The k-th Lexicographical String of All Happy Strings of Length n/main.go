package main

import (
	"math"
)

func main() {
}

func getHappyString(n int, k int) string {
	res, add := "", ""
	c := col(n)
	if k > c {
		return res
	}

	pos, add := first(c, k)
	step := int(c / 3)
	res = add

	for i := 1; i < n; i++ {
		step = int(step / 2) // int((c - pos) / 2)
		l := less(add)
		if k <= pos+step {
			add = l
		} else {
			add = third(add, l)
			pos += step
		}
		res += add
	}

	return res
}

func col(n int) int {
	return int(math.Pow(2, float64(n-1)) * 3)
}

func first(c, k int) (pos int, f string) {
	if k <= int(c/3) {
		return 0, "a"
	}
	if k <= int(c/3)*2 {
		return int(c / 3), "b"
	}
	return int(c/3) * 2, "c"
}

func second(s string, c, k int) string {
	l := less(s)
	if k <= int(c/2) {
		return l
	}
	return third(s, l)
}

func less(s string) string {
	if s == "a" {
		return "b"
	}
	if s == "b" {
		return "a"
	}
	return "a"
}

func third(s1, s2 string) string {
	if s1 != "a" && s2 != "a" {
		return "a"
	}
	if s1 != "b" && s2 != "b" {
		return "b"
	}
	return "c"
}
