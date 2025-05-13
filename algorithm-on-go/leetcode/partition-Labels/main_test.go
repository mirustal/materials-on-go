package partition_Labels

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		excepted []int
	}{
		{
			name:     "first test",
			str:      "ababcbacadefegdehijhklij",
			excepted: []int{9, 7, 8},
		},
		{
			name:     "second test",
			str:      "eccbbbbdec",
			excepted: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partitionLabels(tt.str)
			if !reflect.DeepEqual(result, tt.excepted) {
				t.Errorf("got %v, want %v", result, tt.excepted)
			}
		})
	}
}
