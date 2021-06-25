package offer

import "testing"

func Test_hammingWeight(t *testing.T)  {
	type testCaseOffer15 struct {
		num uint32
		want int
	}

	cases := []testCaseOffer15{
		{num: 11, want: 3},
		{num: 128, want: 1},
		{num: 4294967294, want: 31},
	}

	for _, c := range cases {
		if got := hammingWeight(c.num); got != c.want {
			t.Errorf("test error for input: %v, want: %d, got: %d", c, c.want, got)
		}
	}
}
