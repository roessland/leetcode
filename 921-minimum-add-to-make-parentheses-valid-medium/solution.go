package leetcode926

func minAddToMakeValid(S string) int {
	mustAdd := 0
	numOpen := 0
	for _, r := range S {
		if r == '(' {
			numOpen++
		}
		if r == ')' {
			if numOpen == 0 {
				mustAdd++
				numOpen++
			}
			numOpen--
		}
	}
	return mustAdd + numOpen
}
