package leetcode

import "testing"

func Test_readBinaryWatch(t *testing.T)  {
	type testCase401 struct {
		turnedOn int
		want []string
	}

	cases := []testCase401{
		{turnedOn: 1, want: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
		{turnedOn: 9, want: []string{}},
	}
	for _, c := range cases	{
		if got := readBinaryWatch(c.turnedOn); !comp(got, c.want) {
			t.Errorf("test error for input: %v, want: %v, got: %v", c, c.want, got)
		}
	}
}

func comp(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
