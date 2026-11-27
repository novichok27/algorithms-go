package arrays

import (
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int // элементы, которые должны остаться (порядок не важен)
	}{
		{"basic", []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"all removed", []int{3, 3}, 3, []int{}},
		{"mixed", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement(tt.nums, tt.val)

			if k != len(tt.want) {
				t.Fatalf("k = %d, want %d", k, len(tt.want))
			}

			got := slices.Clone(tt.nums[:k])
			want := slices.Clone(tt.want)
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("nums[:k] = %v, want %v (any order)", got, want)
			}
		})
	}
}
