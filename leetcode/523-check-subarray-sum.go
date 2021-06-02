package leetcode

func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	if l == 0 || k == 0 {
		return false
	}

	m := map[int]int{0: -1}
	sum := 0
	for i,v := range nums {
		sum += v
		mod := sum % k
		if index, ok := m[mod]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			m[mod] = i
		}
	}

	return false
}
