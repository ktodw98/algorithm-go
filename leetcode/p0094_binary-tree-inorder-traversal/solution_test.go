package p0094

import (
	"slices"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want []int
	}{
		{"example1", []int{1, null, 2, 3}, []int{1, 3, 2}},
		{"example2", []int{1, 2, 3, 4, 5, null, 8, null, null, 6, 7, 9}, []int{4, 2, 6, 5, 7, 1, 3, 9, 8}},
		{"example3", []int{}, []int{}},
		{"example4", []int{1}, []int{1}},
		{"example5", []int{1, null, 0, 3}, []int{1, 3, 0}},
	}

	for _, tt := range tests {
		got := solve(SliceToLevelTree(tt.root))
		if !slices.Equal(got, tt.want) {
			t.Errorf("%v) %v => %v (Expected: %v)", tt.name, tt.root, got, tt.want)
		}
	}
}
