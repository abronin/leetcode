package main

import "fmt"

func main() {
	fmt.Println()
}

func reformat(s string) string {
	if len(s) < 2 {
		return s
	}
	chars, digits := []rune{}, []rune{}
	for _, v := range s {
		if isDigit(v) {
			digits = append(digits, v)
		} else {
			chars = append(chars, v)
		}
	}

	l := len(chars) - len(digits)
	if l < -1 || l > 1 {
		return ""
	}

	res := ""
	for i, v := range chars {
		res += string(v)
		if i < len(digits) {
			res += string(digits[i])
		}
	}
	if len(digits) > len(chars) {
		res = string(digits[len(digits)-1]) + res
	}

	return res
}

func isDigit(c rune) bool {
	return byte(c) >= byte('0') && byte(c) <= byte('9')
}
