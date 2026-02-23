package p0070

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"example1", 2, 2},
		{"example2", 3, 3},
		{"example3", 1, 1},
		{"example4", 9, 55},
	}

	for _, tt := range tests {
		got := solve(tt.n)
		if got != tt.want {
			t.Errorf("%v) a_%v = %v (Expected %v)", tt.name, tt.n, got, tt.want)
		}
	}
}
