package main

func main() {
	isValid("]")
}

func isValid(s string) bool {
	slice := []rune(s)
	stack := make([]rune, 0)
	for _, v := range slice {
		if isOpen(v) {
			stack = append(stack, v)
		} else {
			if !validPair(stack, v) {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isOpen(r rune) bool {
	return (r == 40 || r == 91 || r == 123)
}

func validPair(stack []rune, closer rune) bool {
	if len(stack) == 0 {
		return false
	}

	opener := stack[len(stack)-1]
	if (opener == 40 && closer != 41) || (opener == 91 && closer != 93) || (opener == 123 && closer != 125) {
		return false
	}

	return true
}
