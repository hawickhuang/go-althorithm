package offer

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}

	return -1
}

func findRepeatNumber2(nums []int) int {
	length := len(nums)

	i := 0
	for i < length {
		if i == nums[i] {
			i += 1
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
