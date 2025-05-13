package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	lenS, lenP := len(s), len(p)

	if lenS < lenP {
		return nil
	}

	var aq, win [26]int

	for i := range p {
		aq[p[i]-'a']++
		win[s[i]-'a']++
	}

	var res []int

	for i := 0; i < lenS-lenP+1; i++ {
		if aq == win {
			res = append(res, i)
		}

		if i+lenP < lenS {
			win[s[i]-'a']--
			win[s[i+lenP]-'a']++
		}
	}

	return res
}
