package leetcode926

import "testing"
import "github.com/stretchr/testify/assert"

func TestNumFlips(t *testing.T) {
	assert.Equal(t, 0, NumFlips(""))
	assert.Equal(t, 0, NumFlips("1"))
	assert.Equal(t, 0, NumFlips("0"))
	assert.Equal(t, 0, NumFlips("01"))
	assert.Equal(t, 0, NumFlips("01111111111111"))
	assert.Equal(t, 1, NumFlips("11011"))
	assert.Equal(t, 1, NumFlips("01101111111111"))
	assert.Equal(t, 5, NumFlips("011011111111110000"))
}
