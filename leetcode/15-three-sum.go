package leetcode

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, v := range nums {
		if v > 0 {
			return
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i+1
		right := len(nums) - 1

		for right > left {
			if v + nums[left] + nums[right] > 0 {
				right--
			} else if v + nums[left] + nums[right] < 0 {
				left++
			} else {
				ans = append(ans, []int{v, nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				right--
				left++
			}
		}
	}
	return
}
