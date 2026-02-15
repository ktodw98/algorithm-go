package p0027

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantNums []int
		wantVal  int
	}{
		{"example1", []int{3, 2, 2, 3}, 3, []int{2, 2}, 2},
		{"example2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}, 5},
		{"example3", []int{}, 0, []int{}, 0},
		{"example4", []int{1, 1, 1}, 1, []int{}, 0},
	}

	for _, tt := range tests {
		got := solve2(tt.nums, tt.val)
		ans := []int{}

		if got > 0 {
			ans = tt.nums[:got]
		}

		if got != tt.wantVal || !sameElem(ans, tt.wantNums) {
			t.Errorf("(%v) nums: %v / remove %v \n expected %v, but %v / %v, wantnums: %v", tt.name, tt.nums, tt.val, tt.wantVal, got, ans, tt.wantNums)
		}
	}
}

func sameElem(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for _, elem := range a {
		if mapA[elem] == 0 {
			mapA[elem] = 1
		} else {
			mapA[elem] = mapA[elem] + 1
		}
	}

	for _, elem := range b {
		if mapB[elem] == 0 {
			mapB[elem] = 1
		} else {
			mapB[elem] = mapB[elem] + 1
		}
	}

	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}

	return true
}
