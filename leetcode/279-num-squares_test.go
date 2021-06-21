package leetcode

import "testing"

func Test_numSquares(t *testing.T) {
	type testCase279 struct {
		n int
		want int
	}

	cases := []testCase279{
		{n: 12, want: 3},
		{n: 13, want: 2},
		{n: 1, want: 1},
	}

	for _, c := range cases {
		if got := numSquares2(c.n); got != c.want {
			t.Errorf("test error for input: %v, got: %d, want: %d", c, got, c.want)
		}
	}
}
