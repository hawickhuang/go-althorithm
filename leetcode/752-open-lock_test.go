package leetcode

import "testing"

func Test_openLock(t *testing.T) {
	type case752 struct {
		deadends []string
		target string
		want int
	}
	cases := []case752{
		{deadends: []string{"0201", "0101", "0102", "1212", "2002"}, target: "0202", want: 6},
		{deadends: []string{"8888"}, target: "0009", want: 1},
		{deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, target: "8888", want: -1},
		{deadends: []string{"0000"}, target: "8888", want: -1},
	}
	for _, c := range cases {
		if got := openLock(c.deadends, c.target); got != c.want {
			t.Errorf("test error for input: %v, want: %d, got: %d", c, c.want, got)
		}
	}
}
