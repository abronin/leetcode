package main

import "fmt"

func main() {
	fmt.Println()
}

func freqAlphabets(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i+2] == '#' {
			res += twosymbols(s[i], s[i+1])
			i += 2
			continue
		}
		res += onesymbol(s[i])
	}
	return res
}

func twosymbols(a, b byte) string {
	res := string(byte('j') + (a-byte('0'))*10 + (b - byte('0')) - 10)
	return res
}

func onesymbol(a byte) string {
	res := string(byte('a') + a - byte('1'))
	return res
}
