package two_pointers

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func IsPalindrome(s string) bool {
	s = nonAlphanumericRegex.ReplaceAllString(strings.ToLower(s), "")
	if s == "" {
		return true
	}
	var j = len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
