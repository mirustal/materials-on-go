package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var sb strings.Builder

	for _, r := range s {
		if (unicode.IsLetter(r)) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}

	buildString := strings.ToLower(sb.String())

	l, r := 0, len(buildString)-1

	for l < r {
		if buildString[l] != buildString[r] {
			return false
		}
		l++
		r--
	}

	return true
}
