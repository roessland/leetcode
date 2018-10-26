package leetcode771

func numJewelsInStones(J string, S string) int {
	isJewel := make(map[rune]bool)
	for _, r := range J {
		isJewel[r] = true
	}

	numJewels := 0
	for _, r := range S {
		if isJewel[r] {
			numJewels++
		}
	}
	return numJewels
}
