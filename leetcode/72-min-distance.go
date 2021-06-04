package leetcode

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	if n*m == 0 {
		return n + m
	}

	dp := make([][]int, m+1)
	for i:=0;i<m+1;i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j:=0;j<n+1;j++{
		dp[0][j] = j
	}

	for i:=1;i<m+1;i++{
		for j:=1;j<n+1;j++{
			left := dp[i-1][j] + 1
			up := dp[i][j-1] + 1
			leftUp := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftUp += 1;
			}
			dp[i][j] = min(left, min(up, leftUp))
		}
	}
	return dp[m][n]
}
