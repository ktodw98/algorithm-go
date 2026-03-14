package p0002

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{"example1", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"example2", []int{0}, []int{0}, []int{0}},
		{"example3", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, tt := range tests {
		got := ListnodeToArray(solve(ArrayToListnode(tt.l1), ArrayToListnode(tt.l2)))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v) %v plus %v => %v (Expected %v)", tt.name, tt.l1, tt.l2, got, tt.want)
		}
	}
}
