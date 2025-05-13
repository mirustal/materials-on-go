package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	lenNums := len(nums)

	var result = make([]int, lenNums)
	leftPtr, rightPtr := 0, lenNums-1

	for i := lenNums - 1; i >= 0; i-- {
		var square int

		if abs(nums[leftPtr]) > abs(nums[rightPtr]) {
			square = nums[leftPtr] * nums[leftPtr]
			result[i] = square
			leftPtr++
		} else {
			square = nums[rightPtr] * nums[rightPtr]
			result[i] = square
			rightPtr--
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
