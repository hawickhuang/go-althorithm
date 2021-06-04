package leetcode

import "testing"

type testCase64 struct {
	args [][]int
	want int
}

func Test_minPathSum(t *testing.T)  {
	cases := []testCase64{
		{args: [][]int{{1,3,1},{1,5,1},{4,2,1}}, want: 7},
		{args: [][]int{{1,2,3}, {4,5,6}},want: 12},
	}

	for _,c  := range cases {
		if got := minPathSum(c.args); got != c.want {
			t.Errorf("test error with input %v, got: %d, want: %d", c.args, got, c.want)
		}
	}
}
