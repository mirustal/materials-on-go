package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	l, r := 0, n-1

	for i := n - 1; i >= 0; i-- {
		square := 0

		if Abs(nums[l]) < Abs(nums[r]) {
			square = nums[r]
			r--
		} else {
			square = nums[l]
			l++
		}

		result[i] = square * square
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
