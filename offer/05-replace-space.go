package offer

func replaceSpace(s string) string {
	cntOfSpace := 0
	for i := range s {
		if s[i] == ' ' {
			cntOfSpace += 1
		}
	}

	oriLen := len(s)
	newLen := len(s) + 2 * cntOfSpace
	ns := make([]byte, newLen)

	for i := oriLen - 1; i >= 0; i -= 1 {
		if s[i] == ' ' {
			ns[newLen-1] = '0'
			ns[newLen-2] = '2'
			ns[newLen-3] = '%'
			newLen -= 3
		} else {
			ns[newLen-1] = s[i]
			newLen -= 1
		}
	}

	return string(ns)
}
