package p0058

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example1", "Hell World", 5},
		{"Example2", "   fly me   to   the moon  ", 4},
		{"Example3", "luffy is still joyboy", 6},
		{"Example4", "a", 1},
		{"Example5", " ", 0},
		{"Example6", "", 0},
	}

	for _, tt := range tests {
		got := solve(tt.input)
		if got != tt.want {
			t.Errorf("%v) %v: Expected %v, but %v", tt.name, tt.input, tt.want, got)
		}
	}
}
