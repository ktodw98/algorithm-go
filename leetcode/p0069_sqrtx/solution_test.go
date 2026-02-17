package p0069

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"example1", 4, 2},
		{"example2", 1, 1},
		{"example3", 0, 0},
		{"example4", 5, 2},
		{"example5", 20, 4},
	}

	for _, tt := range tests {
		got := solve(tt.input)
		if got != tt.output {
			t.Errorf("%v) input: %v. Expected %v, but %v", tt.name, tt.input, tt.output, got)
		}
	}
}
