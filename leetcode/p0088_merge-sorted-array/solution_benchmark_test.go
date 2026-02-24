package p0088

import "testing"

var BenchmarkM int = 100
var BenchmarkN int = 100

var BenchmarkNums1 []int = func() []int {
	a := make([]int, BenchmarkM+BenchmarkN)
	for i := 0; i < BenchmarkM; i++ {
		a[i] = i
	}

	return a
}()

var BenchmarkNums2 []int = func() []int {
	a := make([]int, BenchmarkN)
	for i := 0; i < BenchmarkN; i++ {
		a = append(a, i)
	}
	return a
}()

func BenchmarkSolve(b *testing.B) {
	nums1 := make([]int, BenchmarkM+BenchmarkN)
	for i := 0; i < b.N; i++ {
		copy(nums1, BenchmarkNums1)
		solve(nums1, BenchmarkM, BenchmarkNums2, BenchmarkN)
	}
}
