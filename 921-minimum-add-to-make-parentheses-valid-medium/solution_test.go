package leetcode926

import "testing"
import "github.com/stretchr/testify/assert"

func TestMinAddToMakeValid(t *testing.T) {
	assert.Equal(t, 4, minAddToMakeValid("()))(("))
}
