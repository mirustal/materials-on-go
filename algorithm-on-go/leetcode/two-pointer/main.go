package two_pointer

func maxArea(height []int) int {
	leftPtr, rightPtr := 0, len(height)-1
	maxArea := -9999

	for leftPtr < rightPtr {
		area := (rightPtr - leftPtr) * min(height[leftPtr], height[rightPtr])

		if area > maxArea {
			maxArea = area
		}

		if height[leftPtr] > height[rightPtr] {
			rightPtr--
		} else {
			leftPtr++
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
