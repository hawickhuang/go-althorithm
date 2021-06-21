package leetcode

import "testing"

func Test_lastStoneWeightII(t *testing.T) {
	type testCase1049 struct {
		stones []int
		want int
	}
	cases := []testCase1049{
		{stones: []int{2,7,4,1,8,1}, want: 1},
		{stones: []int{31,26,33,21,40}, want: 5},
		{stones: []int{1,2}, want: 1},
	}

	for _, c := range cases {
		if got := lastStoneWeightII(c.stones); got != c.want {
			t.Errorf("test error for input: %v, got: %d, want: %d", c, got, c.want)
		}
	}
}
