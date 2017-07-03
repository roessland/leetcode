package leetcode9

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var i int
	for i = 0; i < x; {
		i, x = 10*i+x%10, x/10
	}
	return i == x || i/10 == x
}

func isPalindromeString(x int) bool {
	nums := strconv.Itoa(x)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}
