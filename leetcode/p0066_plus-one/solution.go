package p0066

func solve(digits []int) []int {
	// Transform the slice backward to use append
	// Use var raise to hold carry
	res := make([]int, 0, len(digits)+1)
	raise := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + raise
		if tmp < 10 {
			res = append(res, tmp)
			raise = 0
		} else {
			res = append(res, tmp%10)
			raise = 1
		}
	}

	if raise == 1 {
		res = append(res, 1)
	}

	output := make([]int, 0, len(res))
	for i := len(res) - 1; i >= 0; i-- {
		output = append(output, res[i])
	}
	return output
}
