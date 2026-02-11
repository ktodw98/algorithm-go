package p0014

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{"example1", []string{"flower", "flow", "flight"}, "fl"},
		{"example2", []string{"dog", "racecar", "car"}, ""},
		{"example3", []string{"follower", "flow", "folight"}, "f"},
		{"example4", []string{"follower", "", "folight"}, ""},
	}

	for _, tt := range tests {
		got := solve(tt.input)
		if got != tt.want {
			t.Errorf("%v (%v): expected %v, but %v", tt.name, tt.input, tt.want, got)
		}
	}
}
