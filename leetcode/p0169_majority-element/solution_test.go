package p0169

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{3, 2, 3}, 3},
		{"example1", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"example1", []int{3}, 3},
	}

	for _, tt := range tests {
		got := solve(tt.nums)
		if got != tt.want {
			t.Errorf("%v) majority num of %v is %v (Expected: %v)", tt.name, tt.nums, got, tt.want)
		}
	}
}
