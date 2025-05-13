package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	answer := make([]string, 0)
	l := 0
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i+1]-nums[i] != 1 {
			if l == i {
				answer = append(answer, fmt.Sprintf("%d", nums[i]))
			} else {
				answer = append(answer, fmt.Sprintf("%d->%d", nums[l], nums[i]))
			}
			l = i + 1
		}
	}

	return answer
}
