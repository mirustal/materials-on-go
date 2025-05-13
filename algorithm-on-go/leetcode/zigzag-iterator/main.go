package zigzag_iterator

func myAlgo(arr []int, arr2 []int) []int {
	answer := make([]int, 0)
	lenArr := 0
	if len(arr) > len(arr2) {
		lenArr = len(arr)
	} else {
		lenArr = len(arr2)
	}

	for i := 0; i < lenArr; i++ {
		if i < len(arr) {
			answer = append(answer, arr[i])
		}
		if i < len(arr2) {
			answer = append(answer, arr2[i])
		}
	}

	return answer
}
