package leetcode

func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)
	if ls == 0 {
		if lp == 0 || p[0] == '*' {
			return true
		}
		return false
	}

	dp := make([][]bool, ls+1)
	for i:= 0;i<=ls;i++ {
		dp[i] = make([]bool, lp+1)
	}
	dp[0][0] = true
	for j:=1;j<=lp;j++{
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i<=ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1]=='?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}

	}

	return dp[ls][lp]
}
