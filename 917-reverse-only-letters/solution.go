package leetcode917

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func reverseOnlyLetters(S string) string {
	s := []byte(S)
	i := 0
	j := len(s) - 1
	for {
		for i < len(s) && !isLetter(s[i]) {
			i++
		}
		for 0 <= j && !isLetter(s[j]) {
			j--
		}
		if i < len(s) && 0 <= j && i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		} else {
			break
		}
	}
	return string(s)
}
