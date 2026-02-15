package p0027

import "testing"

var benchVal = 42

// 스왑 최대: 절반이 val, 교대 배치
func makeHalfMatch() []int {
	nums := make([]int, 1000000)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = benchVal
		} else {
			nums[i] = i
		}
	}
	return nums
}

// 스왑 최소: val이 맨 끝에 1개만
func makeFewMatch() []int {
	nums := make([]int, 1000000)
	for i := range nums {
		nums[i] = i
	}
	nums[99999] = benchVal
	return nums
}

func BenchmarkSolve_HalfMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := makeHalfMatch()
		solve(nums, benchVal)
	}
}

func BenchmarkSolve_FewMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := makeFewMatch()
		solve(nums, benchVal)
	}
}

func BenchmarkSolve2_HalfMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := makeHalfMatch()
		solve2(nums, benchVal)
	}
}

func BenchmarkSolve2_FewMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := makeFewMatch()
		solve2(nums, benchVal)
	}
}

// BenchmarkSolve_HalfMatch-10                 1094           1092886 ns/op         8003585 B/op          1 allocs/op
// BenchmarkSolve_FewMatch-10                  1782            660139 ns/op         8003586 B/op          1 allocs/op
// BenchmarkSolve2_HalfMatch-10                1646            716586 ns/op         8003585 B/op          1 allocs/op
// BenchmarkSolve2_FewMatch-10                 1778            661030 ns/op         8003585 B/op          1 allocs/op
