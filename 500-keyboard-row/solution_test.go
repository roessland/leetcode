package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	assert.Equal(t, []string{"Alaska", "Dad"}, findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	assert.Equal(t, []string{"Alaska", "Dad"}, findWordsRegex([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
