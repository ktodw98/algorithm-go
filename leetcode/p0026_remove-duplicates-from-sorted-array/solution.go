package p0026

var MAXINT = 101

func solve(nums []int) int {
	already := make(map[int]bool)
	j := 0

	for i := 0; i < len(nums); i++ {
		if already[nums[i]] {
			nums[i] = MAXINT
		} else {
			already[nums[i]] = true
			if j != i {
				nums[j] = nums[i]
				nums[i] = MAXINT
			}
			j = j + 1
		}
	}
	sort(nums)
	return j
}

func sort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
