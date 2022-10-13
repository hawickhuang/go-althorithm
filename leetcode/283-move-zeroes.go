package leetcode

func moveZeroes(nums []int)  {
	length := len(nums)

	slow, fast := 0, 0

	for fast < length {
		if nums[fast] == 0 {
			fast++
		} else {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
			fast++
		}
	}
}
