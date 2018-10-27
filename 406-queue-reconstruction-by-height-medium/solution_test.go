package leetcode406

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"

func TestSortPeople(t *testing.T) {
	in := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	sortPeople(in)
	exp := [][]int{{7, 0}, {7, 1}, {6, 1}, {5, 0}, {5, 2}, {4, 4}}
	assert.Equal(t, exp[4][0], in[4][0])
	assert.Equal(t, exp[4][1], in[4][1])
}

func TestReconstructQueue(t *testing.T) {
	in := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	exp := [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}
	out := reconstructQueue(in)
	fmt.Printf("%v\n", out)
	assert.Equal(t, exp[0][0], out[0][0])
	assert.Equal(t, exp[3][1], out[3][1])
}
