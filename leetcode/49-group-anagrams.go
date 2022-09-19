package leetcode

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}

	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c-'a']++
		}
		mp[key] = append(mp[key], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
