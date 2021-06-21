package leetcode

func stoneGame(piles []int) bool {
	l := len(piles)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		dp[i][i] = piles[i]
	}

	for i := l-2;i>=0;i-- {
		for j := i+1;j<l;j++ {
			dp[i][j] = max(piles[i] - dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][l-1] > 0
}
