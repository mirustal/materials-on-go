package generate_parentheses

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(s string, open, close int)

	backtrack = func(s string, open, close int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}
		if open < n {
			backtrack(s+"(", open+1, close)
		}
		if close < open {
			backtrack(s+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return result
}
