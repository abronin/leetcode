package main

func main() {

}

func isLongPressedName(name string, typed string) bool {
	index := -1
	for _, c := range name {
		index = findChar(typed, c, index+1)
		if index == -1 {
			return false
		}
	}
	return true
}

func findChar(typed string, c rune, index int) int {
	for i := index; i < len(typed); i++ {
		if typed[i] == byte(c) {
			return i
		}
	}
	return -1
}
