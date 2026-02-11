package p0009

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"example1", 121, true},
		{"example2", -121, false},
		{"example3", 10, false},
		{"example4", 0, true},
	}

	for _, tt := range tests {
		got := solve2(tt.x)
		if got != tt.want {
			t.Errorf("%v (%v): expected %v, but %v", tt.name, tt.x, tt.want, got)
		}
	}
}
