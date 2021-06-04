package leetcode

func findMaxLength(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}

	counter := 0
	m := map[int]int{0:-1}
	ans := 0
	for i, num := range nums{
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if pre, ok := m[counter]; ok {
			if i-pre > ans{
				ans = i-pre
			}
		} else {
			m[counter] = i
		}
	}
	return ans
}
