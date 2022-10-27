package leetcode

import (
	"fmt"
	"sort"
)

func comp2DSlice(s1, s2 [][]string) bool {
	if len(s1) != len(s2) {
		return false
	}

	sort.Slice(s1, func(i, j int) bool {
		return len(s1[i]) < len(s1[j])
	})
	sort.Slice(s2, func(i, j int) bool {
		return len(s2[i]) <len(s2[j])
	})

	for i, v := range s1 {
		if !compSlice(v, s2[i]) {
			return false
		}
	}
	return true
}

func compSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	for i, v := range s1 {
		if s2[i] != v {
			fmt.Println("s2[i]: ", s2[i], "; v: ", v)
			return false
		}
	}

	return true
}