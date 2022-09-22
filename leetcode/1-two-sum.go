package leetcode

func twoSum(nums []int, target int) []int {
	record := make(map[int]int, 0)

	for i, n := range nums {
		if v, ok := record[n]; ok {
			return []int{v, i}
		}
		record[target-n] = i
	}

	return []int{}
}
