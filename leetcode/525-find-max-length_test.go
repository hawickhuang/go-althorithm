package leetcode

import "testing"

type testCase525 struct {
	args []int
	want int
}

func Test_findMaxLength(t *testing.T) {
	cases := []testCase525{
		{args: []int{0}, want: 0},
		{args: []int{1}, want: 0},
		{args: []int{}, want: 0},
		{args: []int{0,1}, want: 2},
		{args: []int{0,1,0}, want: 2},
		{args: []int{0,1,0,1,1,0}, want: 6},
		{args: []int{0,0,0,1,1,1,0}, want: 6},
		{args: []int{1,1,1,1,1,1,1,0,0,0,0,1,1,0,1,0,0,1,1,1,1,1,1,1,1,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,1,1,0,0,0,0,1,0,0,1,1,1,1,1,0,0,1,0,1,1,0,0,0,1,0,0,0,1,1,1,0,1,1,0,1,0,0,1,1,0,1,0,0,1,1,1,0,0,1,0,1,1,1,0,0,1,0,1,1}, want: 94},
	}

	for i := range cases {
		if got := findMaxLength(cases[i].args); got != cases[i].want {
			t.Errorf("test error for input %v, got: %d, want: %d", cases[i].args, got, cases[i].want)
		}
	}
}