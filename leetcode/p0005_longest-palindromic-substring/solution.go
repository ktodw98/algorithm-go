package p0005

func solve(s string) string {
	/*Brute Force with nested loop*/
	maxLen := 0
	maxStart := 0
	for i := 0; i < len(s); i++ {
		j := len(s) - 1
		for i <= j && j-i+1 > maxLen {
			if is_palindrome(s[i:j+1]) && maxLen < j-i+1 {
				maxLen = j - i + 1
				maxStart = i
			}
			j--
		}
	}

	return s[maxStart : maxStart+maxLen]
}

func is_palindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[j] != s[i] {
			return false
		}
	}
	return true
}
