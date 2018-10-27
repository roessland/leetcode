package leetcode338

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"

func TestCountBits(t *testing.T) {
	res := countBits(2)
	exp := []int{0, 1, 1}
	for i := range exp {
		assert.Equal(t, exp[i], res[i], fmt.Sprintf("%v != %v", exp, res))
	}
}
