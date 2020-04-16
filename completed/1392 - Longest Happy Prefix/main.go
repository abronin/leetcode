package main

import "fmt"

func main() {
	s := "aaaaa"
	fmt.Println(s[0:3])
	fmt.Println(longestPrefix(s))
}

func longestPrefix(s string) string {
	longest := ""
	for i := range s {
		if s[0:i+1] == s[len(s)-i-1:] && i != len(s)-1 {
			longest = s[0 : i+1]
		}
	}
	return longest
}
