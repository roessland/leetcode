package leetcode926

/**

abba: a b b a, bb, abba: 6

num palindromes in abba

c(aaa) = c(aa) + c(aa) + c(a

*/

func isPalindrome(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func countSubstringsSlow(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s, i, j) {
				count++
			}
		}
	}
	return count
}

func countSubstrings(s string) int {
	return 0
}
