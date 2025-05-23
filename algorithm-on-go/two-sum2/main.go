package two_sum2

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var sum int

	for left < right {
		sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
