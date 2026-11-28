package arrays

import (
	"slices"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"test1", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"test2", []int{1, 2, 3, 4, 5, 6, 7}, 1, []int{7, 1, 2, 3, 4, 5, 6}},
		{"test3", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)

			if !slices.Equal(tt.nums, tt.want) {
				t.Fatalf("got = %v want = %v", tt.nums, tt.want)
			}
		})
	}
}
