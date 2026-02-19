package p0035

func solve(nums []int, target int) int {
	i, j := 0, len(nums)-1
	var idx int = (i + j) / 2

	for i <= j {
		if nums[idx] == target {
			break
		}

		// In Normal Case
		if nums[idx] > target {
			j = idx - 1
		} else {
			i = idx + 1
		}
		idx = (i + j) / 2
	}
	if nums[idx] < target {
		return idx + 1
	}
	return idx
}
