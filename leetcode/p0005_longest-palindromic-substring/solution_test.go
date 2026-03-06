package p0005

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct{
		name string
		s string
		wantSet map[string]bool
	}{
		{"example1", "babad", map[string]bool{"aba":true, "bab":true,}},
		{"example2", "cbbd", map[string]bool{"bb":true,}},
		{"example3", "b", map[string]bool{"b":true,}},
		{"example4", "kskasasa", map[string]bool{"asasa":true,}},
	}

	for _, tt := range tests {
		got := solve(tt.s)
		_, ok := tt.wantSet[got]
		if !ok {
			t.Errorf("%v) %v => %v (Expected: %v)", tt.name, tt.s, got, tt.wantSet)
		}
	}
}
