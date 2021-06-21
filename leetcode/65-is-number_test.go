package leetcode

import "testing"

func Test_isNumber(t *testing.T) {
	type testCase65 struct {
		s string
		want bool
	}

	cases := []testCase65{
		{s: "0", want: true},
		{s: "e", want: false},
		{s: ".", want: false},
		{s: ".1", want: true},
		{s: "+1", want: true},
		{s: "-1", want: true},
		{s: "+1e2", want: true},
	}

	for _, c := range cases	{
		if got := isNumber(c.s); got != c.want {
			t.Errorf("test error for input: %v, got: %v, want: %v", c, got, c.want)
		}
	}
}
