package leetcode

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type testCase494 struct {
		nums []int
		target int
		want int
	}

	cases := []testCase494{
		{nums: []int{1,1,1,1,1}, target: 3, want: 5},
		{nums: []int{1}, target: 1, want: 1},
	}

	for _, c := range cases {
		if got := findTargetSumWays2(c.nums, c.target); got != c.want {
			t.Errorf("test error for input: %v, want: %d, got: %d", c, c.want, got)
		}
	}
}