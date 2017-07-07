package leetcode14

import "testing"
import "github.com/stretchr/testify/assert"

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "a", LongestCommonPrefix([]string{"a"}))
	assert.Equal(t, "a", LongestCommonPrefix([]string{"a", "a"}))
	assert.Equal(t, "aa", LongestCommonPrefix([]string{"aa", "aa"}))
	assert.Equal(t, "aa", LongestCommonPrefix([]string{"aa", "aab"}))
	assert.Equal(t, "aa", LongestCommonPrefix([]string{"aac", "aab"}))
	assert.Equal(t, "a", LongestCommonPrefix([]string{"aac", "aab", "a"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"", "", ""}))
	assert.Equal(t, "", LongestCommonPrefix([]string{}))
}
