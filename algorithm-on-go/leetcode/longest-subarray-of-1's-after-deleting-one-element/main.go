package longest_subarray_of_1_s_after_deleting_one_element

func longestSubarray(nums []int) int {
	l, r, currVal, numZero := 0, 0, 0, 1
	for r, currVal = range nums {
		if currVal == 0 {
			numZero--
		}

		if numZero < 0 {
			if nums[l] == 0 {
				numZero++
			}
			l++
		}
	}

	return r - l
}
