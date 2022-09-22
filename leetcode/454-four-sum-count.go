package leetcode

func fourSumCount(nums1, nums2, nums3,nums4 []int) int {
	count := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			count[n1+n2]++
		}
	}

	ant := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if v, ok := count[0-(n3+n4)]; ok {
				ant += v
			}
		}
	}
	return ant
}
