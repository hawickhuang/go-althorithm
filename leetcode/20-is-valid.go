package leetcode

func isValid(s string) bool {
	vm := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}

	stack := make([]string, 0)
	for _, c := range s {
		if len(stack) > 0 && vm[string(c)] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(c))
		}
	}
	return len(stack) == 0
}
