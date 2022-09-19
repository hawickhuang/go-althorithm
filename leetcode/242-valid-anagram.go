package leetcode

func validAnagram(s, t string) bool {
	var record [26]int

	for _, c := range s {
		record[c-'a']++
	}
	for _, c := range t {
		record[c-'a']--
	}

	for _, cnt := range record {
		if cnt != 0 {
			return false
		}
	}


	return true
}
