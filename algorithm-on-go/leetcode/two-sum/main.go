package two_sum

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums)-1)
	for i, num := range nums {
		if j, ok := mp[target-num]; ok {
			return []int{j, i}
		} else {
			mp[num] = i
		}
	}
	return nil
}
