package leetcode

import "testing"

func Test_maxPoints(t *testing.T)  {
	type case149 struct {
		points [][]int
		want int
	}

	cases := []case149{
		{points: [][]int{{1,1}, {2, 2}, {3,3}}, want: 3},
		{points: [][]int{{1,1}, {3,2}, {5,3}, {4,1}, {2,3}, {1, 4}}, want: 4},
	}

	for _, c := range cases {
		if got := maxPoints(c.points); got != c.want {
			t.Errorf("test error for input: %v, want: %d, got: %d", c, c.want, got)
		}
	}
}
