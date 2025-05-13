package squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

func testSortedSquares(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		excepted []int
	}{
		{
			name:     "example 1",
			input:    []int{-4, -1, 0, 3, 10},
			excepted: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "example 2",
			input:    []int{-7, -3, 2, 3, 11},
			excepted: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedSquares(tt.input)
			if !reflect.DeepEqual(result, tt.excepted) {
				t.Errorf("sortedSquares() = %v, wanted %v", result, tt.excepted)
			}
		})
	}
}
