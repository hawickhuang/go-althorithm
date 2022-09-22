package leetcode

func commonChars(words []string) (ans []string) {
	if len(words) == 0 {
		return
	}

	var counts [26]int
	for _, c := range words[0] {
		counts[c-'a']++
	}

	for _, word := range words[1:] {
		var wc [26]int
		for _, c := range word {
			wc[c-'a']++
		}

		for i, cnt := range wc {
			if cnt < counts[i] {
				counts[i] = cnt
			}
		}
	}

	for i, count := range counts {
		for j := 0; j < count; j++ {
			ans = append(ans, string(rune('a'+i)))
		}
	}
	return
}
