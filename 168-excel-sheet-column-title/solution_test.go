package leetcode168

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTitle(t *testing.T) {
	assert.Equal(t, "A", convertToTitle(1), "1")
	assert.Equal(t, "Z", convertToTitle(26), "26")
	assert.Equal(t, "AA", convertToTitle(27), "27")
	assert.Equal(t, "AB", convertToTitle(28), "28")
	assert.Equal(t, "AZ", convertToTitle(52), "52")
	assert.Equal(t, "BA", convertToTitle(53), "53")
	assert.Equal(t, "BB", convertToTitle(54), "54")
	assert.Equal(t, "ZY", convertToTitle(701), "701")
	assert.Equal(t, "XQDW", convertToTitle(433443), "433443")
}
