package leetcode14

func get(str string, pos int) (rune, bool) {
	if pos < len(str) {
		return rune(str[pos]), true
	}
	return rune('_'), false
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	common := []rune{}
	for j := 0; ; j++ {
		var topChar rune
		foundTopChar := false
		for i := 0; i < len(strs); i++ {
			ch, ok := get(strs[i], j)
			if !ok {
				return string(common)
			}
			if !foundTopChar {
				topChar = ch
				foundTopChar = true
			} else {
				if topChar != ch {
					return string(common)
				}
			}
		}
		common = append(common, topChar)
	}
	return ""
}
