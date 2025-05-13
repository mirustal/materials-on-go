package permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var win, aq [26]int
	for i := range len(s1) {
		win[s1[i]-'a']++
		aq[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		if win == aq {
			return true
		}

		if i+len(s1) < len(s2) {
			aq[s2[i]-'a']--
			aq[s2[i+len(s1)]-'a']++
		}
	}

	return false
}
