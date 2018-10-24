package leetcode926

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func NumFlips(S string) int {
	// precompute remaining zero count
	count := make([]int, len(S))
	c := 0
	for i := len(count) - 1; i >= 0; i-- {
		if S[i] == '0' {
			c++
		}
		count[i] = c
	}
	var N func(int, int) int
	N = func(i int, total int) int {
		if i == len(S) {
			return total
		}
		if S[i] == '0' {
			return N(i+1, total)
		} else {
			return Min(total+count[i], N(i+1, total+1))
		}
	}
	return N(0, 0)
}
