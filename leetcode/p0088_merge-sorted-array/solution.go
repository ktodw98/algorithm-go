package p0088

func solve(nums1 []int, m int, nums2 []int, n int) {
	/*Fill Backward with two pointer*/
	i, j, current := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[current] = nums1[i]
			i = i - 1
		} else {
			nums1[current] = nums2[j]
			j = j - 1
		}
		current = current - 1
	}
}
