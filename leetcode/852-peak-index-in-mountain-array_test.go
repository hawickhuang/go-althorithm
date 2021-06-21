package leetcode

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type testCase852 struct {
		arr []int
		want int
	}
	cases := []testCase852{
		{arr: []int{0, 1, 0}, want: 1},
		{arr: []int{0,2,1,0}, want: 1},
		{arr: []int{0, 10, 5, 2}, want: 1},
		{arr: []int{3,4,5,1}, want: 2},
		{arr: []int{24,69,100,99,79,78,67,36,26,19}, want: 2},
	}
	for _, c := range cases {
		if got := peakIndexInMountainArray(c.arr); got != c.want {
			t.Errorf("test error for input: %v, got: %d, want: %d", c, got, c.want)
		}
	}
}