package subarray_sum_equals_k

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name        string
		input_array []int
		input_num   int
		excepted    int
	}{
		{
			name:        "example 1",
			input_array: []int{1, 1, 1},
			input_num:   2,
			excepted:    2,
		},
		{
			name:        "example 2",
			input_array: []int{1, 2, 3},
			input_num:   3,
			excepted:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraySum(tt.input_array, tt.input_num)
			if result != tt.excepted {
				t.Errorf("subarraySum() = %d, want %d", result, tt.excepted)
			}
		})
	}
}
