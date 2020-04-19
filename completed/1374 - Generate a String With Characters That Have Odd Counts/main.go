package main

import "fmt"

func main() {
	fmt.Println(generateTheString(4))
	fmt.Println(generateTheString(2))
	fmt.Println(generateTheString(26))
}

func generateTheString(n int) string {
	x := n / 25
	res := xString(x)
	o := n % 25
	if o == 0 {
		return res
	}
	if o%2 == 1 {
		res += mult(o, "z")
	} else {
		res += "y" + mult(o-1, "z")
	}
	return res
}

func xString(x int) string {
	res := ""
	for i := 0; i < x; i++ {
		res += mult(25, string(byte('a'+i)))
	}
	return res
}

func mult(x int, c string) string {
	res := ""
	for i := 0; i < x; i++ {
		res += c
	}
	return res
}
