package arrays

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test1", []int{1, 1, 2}, []int{1, 2}},
		{"test2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{"test3", []int{0, 0, 0, 0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.nums)

			if k != len(tt.want) {
				t.Fatalf("k = %d want = %d", k, len(tt.want))
			}

			got := slices.Clone(tt.nums[:k])
			want := slices.Clone(tt.want)
			if !slices.Equal(got, want) {
				t.Errorf("nums[:k] = %v, want %v", got, want)
			}
		})
	}

}
