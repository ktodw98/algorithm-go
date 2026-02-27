package p0169

func solve(nums []int) int {
	mapper := make(map[int]int)

	maxNum, maxNumCnt := 0, 0
	for _, v := range nums {
		mapper[v] = mapper[v] + 1
		if maxNumCnt < mapper[v] {
			maxNum, maxNumCnt = v, mapper[v]
		}
	}
	return maxNum
}
