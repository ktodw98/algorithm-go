package p0021

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		lst1 []int
		lst2 []int
		want []int
	}{
		{"example1", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"example2", []int{}, []int{}, []int{}},
		{"example3", []int{}, []int{0}, []int{0}},
		{"example4", []int{1, 2}, []int{}, []int{1, 2}},
	}

	for _, tt := range tests {
		fmt.Println(tt.name)
		ln1, ln2 := SliceToListNode(tt.lst1), SliceToListNode(tt.lst2)
		got := ListNodeToSlice(solve(ln1, ln2))
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("%v) %v, %v => Expected %v, but %v", tt.name, tt.lst1, tt.lst2, tt.want, got)
		}
	}
}
