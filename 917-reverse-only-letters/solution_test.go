package leetcode917

import "github.com/stretchr/testify/assert"
import "testing"

func TestIsLetter(t *testing.T) {
	assert.True(t, isLetter('a'))
	assert.False(t, isLetter('-'))
	assert.False(t, isLetter('='))
}

func TestReverseOnlyLetters(t *testing.T) {
	assert.Equal(t, "a-ab", reverseOnlyLetters("b-aa"))
	assert.Equal(t, "dc-ba", reverseOnlyLetters("ab-cd"))
	assert.Equal(t, "=+-1-caa-ba", reverseOnlyLetters("=+-1-aba-ac"))
	assert.Equal(t, "", reverseOnlyLetters(""))
	assert.Equal(t, "-", reverseOnlyLetters("-"))
}
