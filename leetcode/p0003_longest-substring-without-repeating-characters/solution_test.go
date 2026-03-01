package p0003

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example1", "abcabcbb", 3},
		{"example2", "bbbbb", 1},
		{"example3", "pwwkew", 3},
		{"example4", "", 0},
		{"example5", "a", 1},
	}

	for _, tt := range tests {
		got := solve(tt.input)
		if got != tt.want {
			t.Errorf("%v) %v => %v (Expected: %v)", tt.name, tt.input, got, tt.want)
		}
	}
}
