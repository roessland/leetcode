package leetcode9

import (
	"testing"

	"math"

	"github.com/stretchr/testify/assert"
)

// Totally overkill testing
func TestIsPalindrome(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		t.Parallel()
		for x := math.MaxInt32 / 2; x <= math.MaxInt32; x++ {
			assert.Equal(t, isPalindromeString(x), isPalindrome(x), x)
		}
	})

	t.Run("B", func(t *testing.T) {
		t.Parallel()
		for x := 0; x < math.MaxInt32/2; x++ {
			assert.Equal(t, isPalindromeString(x), isPalindrome(x), x)
		}
	})
	t.Run("C", func(t *testing.T) {
		t.Parallel()
		for x := math.MinInt32 / 2; x < 0; x++ {
			assert.Equal(t, isPalindromeString(x), isPalindrome(x), x)
		}
	})

	t.Run("D", func(t *testing.T) {
		t.Parallel()
		for x := math.MinInt32; x < math.MinInt32/2; x++ {
			assert.Equal(t, isPalindromeString(x), isPalindrome(x), x)
		}
	})
}
