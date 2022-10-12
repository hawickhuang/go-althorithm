package leetcode

import "fmt"

func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := getNext(needle)
	fmt.Println( next)

	j := 0
	for i:=0; i<len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		fmt.Println("i: ", i, ";j: ", j, ":", len(next)-1, ";", i-len(next)+1)
		if j == len(next) {
			return i-len(next)+1
		}
	}
	return -1
}

func getNext(needle string) []int {
	l := len(needle)
	ans := make([]int, l)
	j := 0
	ans[0] = j

	for i:=1; i<l; i++ {
		for j>0 && needle[i] != needle[j] {
			j = ans[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		ans[i] = j
	}
	return ans
}