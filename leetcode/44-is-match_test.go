package leetcode

import "testing"

type testCase44 struct {
	args [2]string
	want bool
}

func Test_isMatch(t *testing.T)  {
	cases := []testCase44{
		{args: [2]string{"aa", "a"}, want: false},
		{args: [2]string{"aa", "*"}, want: true},
		{args: [2]string{"ch", "?a"}, want: false},
		{args: [2]string{"adceb", "*a*b"}, want: true},
		{args: [2]string{"acdcb", "a*c?b"}, want: false},
	}

	for i := range cases {
		if got := isMatch(cases[i].args[0], cases[i].args[1]); got != cases[i].want {
			t.Errorf("test error with input %v, want: %v, got: %v", cases[i].args, cases[i].want, got)
		}
	}
}