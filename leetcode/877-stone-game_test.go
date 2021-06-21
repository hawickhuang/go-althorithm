package leetcode

import "testing"

func Test_stoneGame(t *testing.T) {
	type testCase877 struct {
		piles []int
		want bool
	}

	cases := []testCase877{
		{piles: []int{5,3,4, 5}, want: true},
	}

	for _, c := range cases {
		if got := stoneGame(c.piles); got != c.want {
			t.Errorf("test error for input: %v, got: %v, wang: %v", c, got, c.want)
		}
	}
}
