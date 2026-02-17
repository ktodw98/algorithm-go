package p0069

func solve(x int) int {
	lastInt := 0
	for i := 1; i <= x; i++ {
		if i*i > x {
			break
		}
		lastInt = i
	}
	return lastInt
}
