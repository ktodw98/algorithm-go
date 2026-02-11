package p0014

func solve(strs []string) string {
	output := ""
	minLen := len(strs[0])

	for _, str := range strs {
		minLen = min(minLen, len(str))
	}

	i := 0
	for i < minLen {
		std := strs[0][i]
		for _, str := range strs {
			if str[i] != std {
				return output
			}
		}
		output = output + string(std)
		i += 1
	}

	return output
}
