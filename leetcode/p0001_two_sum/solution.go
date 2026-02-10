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
