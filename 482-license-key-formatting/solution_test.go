package leetcode482

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseKeyFormatting(t *testing.T) {
	assert.True(t, true)
	assert.Equal(t, "24A0-R74K", licenseKeyFormatting("2-4A0r7-4k", 4))
	assert.Equal(t, "24-A0R-74K", licenseKeyFormatting("2-4A0r7-4k", 3))
}
