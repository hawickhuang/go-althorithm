package leetcode

func findAnagrams(s string, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	var sCount, pCount [26]int
	for i, c := range p {
		pCount[c-'a']++
		sCount[s[i]-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, c := range s[:sLen-pLen] {
		sCount[c-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
