package leetcode

import "testing"

type testCase72 struct {
	word1 string
	word2 string
	want int
}

func Test_minDistance(t *testing.T)  {
	cases := []testCase72{
		{word1: "horse", word2: "ros", want: 3},
		{word1: "intention", word2: "execution", want: 5},
	}

	for _, c := range cases {
		if got := minDistance(c.word1, c.word2); got != c.want {
			t.Errorf("test error for input: %v, got: %d, want: %d", c, got, c.want)
		}
	}
}