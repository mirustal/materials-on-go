package valid_validrom

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var sb strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}

	builtString := strings.ToLower(sb.String())

	size := len(builtString)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		if builtString[i] != builtString[j] {
			return false
		}
	}
	return true
}
