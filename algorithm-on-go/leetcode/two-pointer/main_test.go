package two_pointer

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "example from leetcode",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "two elements",
			height:   []int{1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.height)
			if result != tt.expected {
				t.Errorf("maxArea() = %d, want %d", result, tt.expected)
			}
		})
	}
}
