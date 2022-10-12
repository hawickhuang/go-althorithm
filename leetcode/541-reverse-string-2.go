package leetcode

func reverseString2(s string, k int) string {
	sb := []byte(s)
	for i:=0; i<len(sb); i+= 2*k {
		if i+k < len(sb) {
			reverse(&sb, i, i+k-1)
			continue
		}
		reverse(&sb, i, len(sb))
	}
	return string(sb)
}


