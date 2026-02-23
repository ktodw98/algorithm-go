package p0083

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"example1", []int{1, 1, 2}, []int{1, 2}},
		{"example2", []int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
		{"example3", []int{}, []int{}},
		{"example4", []int{-1, -1, 0, 1, 2, 2}, []int{-1, 0, 1, 2}},
		{"example5", []int{0, 0, 0, 0, 0}, []int{0}},
		{"example6", []int{0, 1, 1}, []int{0, 1}},
	}

	for _, tt := range tests {
		got := ListNodeToSlice(solve(SliceToListNode(tt.input)))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v) %v => %v (Expected %v)", tt.name, tt.input, got, tt.want)
		}
	}
}
