package leetcode

func isHappy(n int) (ans bool) {
	repeat := make(map[int]struct{}, 0)

	for {
		var sum int
		for {
			if n == 0 {
				break
			}
			sum += (n%10)*(n%10)
			n /= 10
		}
		if sum == 1 {
			ans = true
			break
		} else if _, ok := repeat[sum]; ok{
			ans = false
			break
		}
		repeat[sum] = struct{}{}
		n = sum
	}
	return
}
