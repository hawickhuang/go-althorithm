package leetcode

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}

func numSquares2(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if check4(n) {
		return 4
	}
	for i:=1; i*i <n;i++{
		if isPerfectSquare(n-i*i) {
			return 2
		}
	}

	return 3
}

func isPerfectSquare(n int) bool {
	y := int(math.Sqrt(float64(n)))
	return n == y*y
}

func check4(n int) bool {
	for n%4 == 0 {
		n /= 4
	}
	return n%8 == 7
}
