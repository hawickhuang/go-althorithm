package leetcode

func minPathSum(grid [][]int) int  {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i:= 1; i<n;i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i:= 1; i<m;i++{
		dp[0] = dp[0] + grid[i][0]
		for j:=1;j<n;j++ {
			dp[j] = min(grid[i][j] + dp[j-1], grid[i][j]+dp[j])
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
