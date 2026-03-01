package p0003

func solve(s string) int {
	/*Use Set to track unique numbers and two pointer*/
	if len(s) == 0 {
		return 0
	}

	set := make(map[rune]int)

	i, j := 0, 0
	set[rune(s[i])] = 1
	maxCnt := 1
	for j < len(s)-1 {
		j++
		if set[rune(s[j])] == 1 {
			for i < j {
				if rune(s[i]) == rune(s[j]) {
					set[rune(s[i])] = 0
					i++
					break
				} else {
					set[rune(s[i])] = 0
					i++
				}
			}
		}
		set[rune(s[j])] = 1
		maxCnt = max(maxCnt, j-i+1)
	}
	return maxCnt
}
