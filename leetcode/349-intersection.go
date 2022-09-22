package leetcode

func intersection(nums1, nums2 []int) (ans []int) {
	nLen1 := len(nums1)
	nmap := make(map[int]struct{}, nLen1)
	for _, n1 := range nums1 {
		nmap[n1] = struct{}{}
	}

	has := make(map[int]struct{}, 0)
	for _, n2 := range nums2 {
		if _, ok := has[n2]; ok {
			continue
		}
		has[n2] = struct{}{}
		if _, ok := nmap[n2]; ok {
			ans = append(ans, n2)
		}
	}
	return
}
