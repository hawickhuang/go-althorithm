package leetcode

import "testing"

func Test_change(t *testing.T)  {
	type testCase struct {
		amount int
		coins []int
		want int
	}

	cases := []testCase{
		{amount: 5, coins: []int{1, 2, 5}, want: 4},
		{amount: 3, coins: []int{2}, want: 0},
		{amount: 10, coins: []int{10}, want: 1},
	}

	for _, c := range cases {
		if got := change(c.amount, c.coins); got != c.want {
			t.Errorf("test error for input: %v, got: %d, want: %d", c, got, c.want)
		}
	}
}