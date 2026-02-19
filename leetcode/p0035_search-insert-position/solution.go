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

func solve2(nums []int, target int) int {
	// Refactor to Go-ish code
	i, j := 0, len(nums)

	// (1) Remove redundant if statement replacing "<=" to "<"
	for i < j {
		// (2) prevent overflow with uint
		// (alternative) mid := i + (j-i)/2
		// (3) boost calc speed with bit operation
		mid := int(uint(i+j) >> 1)

		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return i
}
