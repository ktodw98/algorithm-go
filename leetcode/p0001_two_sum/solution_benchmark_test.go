package p0001

import "testing"

// worst case: 정답이 배열 맨 끝에 위치 (O(n²) 전부 탐색)
var benchNums = func() []int {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i
	}
	return nums
}()

var benchTarget = benchNums[99998] + benchNums[99999]

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(benchNums, benchTarget)
	}
}
