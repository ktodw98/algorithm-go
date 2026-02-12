package p0020

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"example1", "()", true},
		{"example2", "()[]{}", true},
		{"example3", "(])", false},
		{"example4", "([])", true},
		{"example5", "([)]", false},
		{"example6", "(", false},
	}

	for _, tt := range tests {
		got := solve(tt.input)
		if got != tt.want {
			t.Errorf("%v: %v => Expected %v, but %v", tt.name, tt.input, tt.want, got)
		}
	}
}
