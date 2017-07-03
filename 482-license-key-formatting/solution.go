package leetcode482

import "strings"

func licenseKeyFormatting(input string, groupLength int) string {
	input = strings.Replace(input, "-", "", -1)
	firstLength := len(input) % groupLength
	var out []string
	if firstLength > 0 {
		out = []string{strings.ToUpper(input[:firstLength])}
	} else {
		out = []string{}
	}

	for i := firstLength; i <= len(input)-groupLength; i += groupLength {
		out = append(out, strings.ToUpper(input[i:i+groupLength]))
	}
	return strings.Join(out, "-")
}
