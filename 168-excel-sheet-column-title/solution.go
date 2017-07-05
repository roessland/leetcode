package leetcode168

func convertToTitle(n int) string {
	title := []rune{}
	for n != 0 {
		title = append([]rune{rune((n-1)%26) + 'A'}, title...)
		n = (n - 1) / 26
	}
	return string(title)
}
