package partition_Labels

func partitionLabels(s string) []int {
	lastChar := make([]int, 26)

	for in, ch := range s {
		lastChar[ch-'a'] = in
	}

	var result []int

	startPoint, endPoint := 0, 0

	for in, ch := range s {
		if lastChar[ch-'a'] > endPoint {
			endPoint = lastChar[ch-'a']
		}

		if in == endPoint {
			result = append(result, endPoint-startPoint+1)
			startPoint = endPoint + 1
		}
	}

	return result
}
