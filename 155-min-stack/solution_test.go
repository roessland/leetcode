package leetcode155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	assert.Equal(t, -3, ms.GetMin())
	ms.Pop()
	assert.Equal(t, 0, ms.Top())
	assert.Equal(t, -2, ms.GetMin())
}
