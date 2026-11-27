package arrays

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test1", []int{3, 2, 3}, 3},
		{"test2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"test3", []int{3, 3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := majorityElement(tt.nums)

			if k != tt.want {
				t.Fatalf("got = %d want = %d", k, tt.want)
			}
		})
	}
}
