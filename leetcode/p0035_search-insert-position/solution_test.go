package p0035

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"example1", []int{1, 3, 5, 6}, 5, 2},
		{"example2", []int{1, 3, 5, 6}, 2, 1},
		{"example3", []int{1, 3, 5, 6}, 7, 4},
		{"example4", []int{0}, 1, 1},
		{"example5", []int{1}, 0, 0},
	}

	for _, tt := range tests {
		got := solve2(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("%v) exact index of %v in %v: Expected %v, but %v", tt.name, tt.target, tt.nums, tt.want, got)
		}
	}
}
