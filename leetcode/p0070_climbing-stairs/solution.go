package p0070

func solve0(n int) int {
	// Time Limit Exceeded
	if n < 0 {
		return 0
	} else if n < 2 {
		return 1
	} else {
		return solve0(n-1) + solve0(n-2)
	}
}

func solve(n int) int {
	// Dynamic Programming
	if n < 0 {
		return 0
	} else if n < 2 {
		return 1
	}
	a := make([]int, n+1)
	a[0], a[1] = 1, 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n]
}
