package is_subsequence

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if s == "" {
		return true
	}
	ptrTarget := 0
	for i := 0; i < len(t); i++ {
		if s[ptrTarget] == t[i] {
			ptrTarget++
			if ptrTarget == len(s) {
				return true
			}
		}
	}

	return false
}
