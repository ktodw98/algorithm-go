package p0058

func solve(s string) int {
	i := len(s) - 1
	cnt := 0

	for i >= 0 && s[i] == ' ' {
		i -= 1
	}
	for j := i; j >= 0; j-- {
		if s[j] != ' ' {
			cnt += 1
		} else {
			return cnt
		}
	}

	return cnt
}
