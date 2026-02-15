package p0028

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"Example1", "sasbutsadsad", "sad", 6},
		{"Example2", "leetcode", "leeto", -1},
		{"Example3", "a", "a", 0},
	}

	for _, tt := range tests {
		got := solve2(tt.haystack, tt.needle)
		if got != tt.want {
			t.Errorf("(%v) the index of '%v' at '%v': Expected %v, but %v", tt.name, tt.needle, tt.haystack, tt.want, got)
		}
	}
}
