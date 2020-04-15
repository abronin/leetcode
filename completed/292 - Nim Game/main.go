package main

import "fmt"

func main() {
	fmt.Println()
}

func canWinNim(n int) bool {
	if n == 0 {
		return true
	}

	return n%4 != 0
}
