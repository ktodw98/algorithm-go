package p0067

import (
	"strconv"
)

func solve(a string, b string) string {
	/*Add backward digit by digit*/
	var ans string

	i, j := len(a)-1, len(b)-1
	nums, carry := 0, 0

	for {
		if i < 0 && j < 0 {
			if carry > 0 {
				ans = strconv.Itoa(carry) + ans
			}
			break
		} else if i >= 0 && j >= 0 {
			a1, _ := strconv.Atoi(string(a[i]))
			b1, _ := strconv.Atoi(string(b[j]))
			nums = a1 + b1 + carry
			i, j = i-1, j-1
		} else if i >= 0 {
			a1, _ := strconv.Atoi(string(a[i]))
			nums = a1 + carry
			i = i - 1
		} else {
			b1, _ := strconv.Atoi(string(b[j]))
			nums = b1 + carry
			j = j - 1
		}

		nums, carry = nums%2, nums/2
		ans = strconv.Itoa(nums) + ans
	}
	return ans
}
