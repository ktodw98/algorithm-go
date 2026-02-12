package p0020

func solve(s string) bool {
	stack := make([]string, 0)

	// classifier whether parenthesis is open or not
	isOpen := make(map[string]bool)
	isOpen["("] = true
	isOpen["{"] = true
	isOpen["["] = true
	isOpen[")"] = false
	isOpen["}"] = false
	isOpen["]"] = false

	pair := make(map[string]string)
	pair["("] = ")"
	pair["{"] = "}"
	pair["["] = "]"

	for i := 0; i < len(s); i = i + 1 {

		if isOpen[string(s[i])] {
			stack = append(stack, string(s[i]))
		} else if len(stack) == 0 || pair[stack[len(stack)-1:][0]] != string(s[i]) {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
