package container_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var maxArea int
	for left < right {
		area := calc(height[left], height[right], left, right)
		if area > maxArea {
			maxArea = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func calc(hl, hr, l, r int) int {
	d := r - l

	if hl > hr {
		return d * hr
	}

	return d * hl
}
