package leetcode

import "testing"

func Test_profitableSchemes(t *testing.T) {
	type testCase879 struct {
		n         int
		minProfit int
		group     []int
		profit    []int
		want      int
	}

	cases := []testCase879{
		{n: 5, minProfit: 3, group: []int{2, 2}, profit: []int{2, 3}, want: 2},
		{n: 10, minProfit: 5, group: []int{2, 3, 5}, profit: []int{6, 7, 8}, want: 7},
	}
	for _, c := range cases {
		if got := profitableSchemes(c.n, c.minProfit, c.group, c.profit); got != c.want {
			t.Errorf("test error for input: %v, got: %d, wang: %d", c, got, c.want)
		}
	}
}
