package leetcode

func removeDuplicates(nums []int) int {
	slow, fast := 0, 1

	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow+1
}
