package leetcode926

import "testing"
import "github.com/stretchr/testify/assert"

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("bob", 0, 0))
	assert.True(t, isPalindrome("bob", 0, 2))
	assert.False(t, isPalindrome("bbo", 1, 2))
}

func TestCountSubstrings(t *testing.T) {
	assert.Equal(t, 3, countSubstrings("abc"))
	assert.Equal(t, 6, countSubstrings("aaa"))
	assert.Equal(t, 4, countSubstrings("bob"))
}
