package p0088

import (
	"slices"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"example1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"example2", []int{1}, 1, []int{}, 0, []int{1}},
		{"example3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"example4", []int{}, 0, []int{}, 0, []int{}},
		{"example5", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"example6", []int{2, 2, 2, 0, 0, 0}, 3, []int{2, 2, 2}, 3, []int{2, 2, 2, 2, 2, 2}},
	}

	for _, tt := range tests {
		origin := make([]int, len(tt.nums1))
		copy(origin, tt.nums1)
		solve(tt.nums1, tt.m, tt.nums2, tt.n)
		if !slices.Equal(tt.nums1, tt.want) {
			t.Errorf("%v) solve(%v, %v, %v, %v) => %v (Expected %v)", tt.name, origin, tt.m, tt.nums2, tt.n, tt.nums1, tt.want)
		}
	}
}
