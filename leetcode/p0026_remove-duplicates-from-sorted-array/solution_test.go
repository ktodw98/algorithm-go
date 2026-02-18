package p0026

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		want     int
		wantList []int
	}{
		{"example1", []int{1, 1, 2}, 2, []int{1, 2}},
		{"example2", []int{0, 0, 4, 1, 1, 2, 2, 3, 3, 1}, 5, []int{0, 1, 2, 3, 4}},
		{"example1", []int{1}, 1, []int{1}},
		{"example1", []int{1, -1}, 2, []int{-1, 1}},
	}

	for _, tt := range tests {
		got := solve(tt.nums)
		if got != tt.want || !isEqual(tt.nums, tt.wantList) {
			t.Errorf("=====\n%v) lst: %v. Expected %v. But %v. (%v)\n=====", tt.name, tt.nums, tt.want, got, tt.wantList)
		}
	}
}
func isEqual(a []int, b []int) bool {
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
