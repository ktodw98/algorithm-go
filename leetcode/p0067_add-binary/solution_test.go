package p0067

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"example1", "11", "1", "100"},
		{"example2", "1010", "1011", "10101"},
		{"example3", "0", "0", "0"},
	}

	for _, tt := range tests {
		got := solve(tt.a, tt.b)

		if got != tt.want {
			t.Errorf("%v) %v + %v => Expected %v, but %v", tt.name, tt.a, tt.b, tt.want, got)
		}
	}
}
