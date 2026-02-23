package p0001

func solve(nums []int, target int) []int {
	output := make([]int, 2)

	N := len(nums)
Loop:
	for i := 0; i < N; i = i + 1 {
		for j := i + 1; j < N; j = j + 1 {
			if nums[i]+nums[j] == target {
				output[0] = i
				output[1] = j
				break Loop
			}
		}
	}

	return output
}

func solve2(nums []int, target int) []int {
	/*Use hashtable */
	mapper := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		idx, ok := mapper[comp]
		if ok && idx != i {
			return []int{idx, i}
		}
		mapper[nums[i]] = i
	}

	return nil
}
