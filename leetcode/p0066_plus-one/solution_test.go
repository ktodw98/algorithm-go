package p0066

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"example1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"example2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"example3", []int{9}, []int{1, 0}},
	}

	for _, tt := range tests {
		got := solve(tt.digits)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v) %v. Expected %v, but %v", tt.name, tt.digits, tt.want, got)
		}
	}
}
