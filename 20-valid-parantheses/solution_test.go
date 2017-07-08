package leetcode20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, IsValid("()"))
	assert.Equal(t, true, IsValid("([])"))
	assert.Equal(t, true, IsValid("([a])"))
	assert.Equal(t, true, IsValid("{a{a{a}a}a}"))
	assert.Equal(t, false, IsValid("()["))
	assert.Equal(t, false, IsValid("([[]}])"))
	assert.Equal(t, false, IsValid("([[a])"))
	assert.Equal(t, false, IsValid("{a{[a{a}a}a}"))
}
