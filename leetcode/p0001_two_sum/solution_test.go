package p0001

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	want := []int{0, 1}

	got := solve(nums, target)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("twoSome() = %v, want %v", got, want)
	}
}

func TestSolveWithSlice(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"기본 케이스", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"가운데 값", []int{3, 2, 4}, 6, []int{1, 2}},
		{"같은 숫자", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		got := solve(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v: twoSome() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
