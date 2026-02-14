package p0027

func solve(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] == val {
			j -= 1
		}
		for i < j && nums[i] != val {
			i += 1
		}
		if nums[i] == val && nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
		} else if i+1 == j {
			break
		}
	}

	if nums[i] == val {
		return i
	}
	return i + 1
}
