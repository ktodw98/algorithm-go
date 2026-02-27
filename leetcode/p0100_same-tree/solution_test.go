package p0100

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		p    []int
		q    []int
		want bool
	}{
		{"example1", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"example2", []int{1, 2}, []int{1, null, 2}, false},
		{"example3", []int{1, 2, 1}, []int{1, 1, 2}, false},
		{"example4", []int{1}, []int{1}, true},
		{"example5", []int{}, []int{}, true},
	}

	for _, tt := range tests {
		a, b := SliceToTree(tt.p), SliceToTree(tt.q)

		got := solve(a, b)
		if got != tt.want {
			t.Errorf("%v) %v and %v => %v (Expected: %v)", tt.name, tt.p, tt.q, got, tt.want)
		}
	}
}
