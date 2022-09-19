package leetcode


func canConstruct(ransomNote string, magazine string) bool {
	var record [26]int

	for _, c := range magazine {
		record[c-'a']++
	}
	for _, c := range ransomNote {
		record[c-'a']--
	}

	for _, cnt := range record {
		if cnt < 0 {
			return false
		}
	}

	return true
}