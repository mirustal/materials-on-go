package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)

	left := 0
	ml := 0
	maxFreq := 0

	for right := 0; right < len(s); right++ {
		char := s[right] - 'A'
		freq[char]++
		maxFreq = max(maxFreq, freq[char])
		if (right-left+1)-maxFreq > k {
			freq[s[left]-'A']--
			left++
		}
		ml = max(ml, right-left+1)
	}

	return ml
}
