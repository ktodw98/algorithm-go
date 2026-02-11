package p0009

import (
	"strconv" // solve2
)

func solve(x int) bool {
	if x < 0 {
		return false
	}

	origin := x

	reverse := 0
	a := x
	for x > 0 {
		a = x / 10
		reverse = reverse*10 + (x % 10)
		x = a
	}

	return reverse == origin
}

func solve2(x int) bool {
	origin := strconv.Itoa(x)

	N := len(origin)
	for i, j := 0, N-1; i <= j; i, j = i+1, j-1 {
		if origin[i] != origin[j] {
			return false
		}
	}
	return true
}
