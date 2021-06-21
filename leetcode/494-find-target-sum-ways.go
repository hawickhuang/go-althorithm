package leetcode

func findTargetSumWays(nums []int, target int) int {
	l := len(nums)
	ans := 0

	var bt func(int, int)
	bt = func(i int, sum int) {
		if i == l {
			if sum == target {
				ans++
			}
			return
		}
		bt(i+1, sum+nums[i])
		bt(i+1, sum-nums[i])
	}
	bt(0, 0)

	return ans
}

func findTargetSumWays2(nums []int, target int) int {
	l := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff % 2 == 1 {
		return 0
	}
	neg := diff / 2

	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1

	for i:=0;i<l;i++{
		for j:=0;j<=neg;j++{
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[l][neg]
}
