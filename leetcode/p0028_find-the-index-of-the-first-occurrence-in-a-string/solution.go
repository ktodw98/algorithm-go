package p0028

import "reflect"

func solve(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if reflect.DeepEqual(haystack[i:i+len(needle)], needle) {
			return i
		}
	}
	return -1
}

func solve2(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func solve3(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		flag := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

func solve4(haystack string, needle string) int {
OuterLoop:
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue OuterLoop
			}
		}
		return i
	}
	return -1
}
