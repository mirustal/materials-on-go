package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	preSum := 0

	res := 0
	for _, num := range nums {
		preSum += num
		if count, exist := mp[preSum-k]; exist {
			res += count
		}

		mp[preSum]++
	}

	return res
}
