package arrays

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test1", []int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}},
		{"test2", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates2(tt.nums)
			if k != len(tt.want) {
				t.Fatalf("k = %d want = %d", k, len(tt.want))
			}

			got := slices.Clone(tt.nums[:k])
			want := slices.Clone(tt.want)
			if !slices.Equal(got, want) {
				t.Fatalf("got = %v want = %v", got, want)
			}
		})
	}
}
