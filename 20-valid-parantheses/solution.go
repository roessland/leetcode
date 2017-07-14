package leetcode20

func push(stack []rune, elem rune) []rune {
	return append(stack, elem)
}

func peek(stack []rune) rune {
	return stack[len(stack)-1]
}

func pop(stack []rune) ([]rune, rune) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

// IsValid checks if a string has a valid parantheses order. Given a string
// containing just the characters '(', ')', '{', '}', '[' and ']', it determines
// if the input string is valid. Works for unicode too.
func IsValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		for _, p := range [][]rune{[]rune{'(', ')'}, []rune{'{', '}'}, []rune{'[', ']'}} {
			if ch == p[0] {
				stack = push(stack, rune(p[0]))
			} else if ch == p[1] {
				if len(stack) == 0 || peek(stack) != p[0] {
					return false
				}
				stack, _ = pop(stack)
			}
		}
	}
	return len(stack) == 0
}

var isValid = IsValid
