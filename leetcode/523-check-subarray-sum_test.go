package leetcode

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	want := true
	got := checkSubarraySum([]int{23, 2, 4, 6, 6}, 7)
	if got != want {
		t.Errorf("test checkSubarraySum error, got: %v, want: %v", got, want)
	}

	want = false
	got = checkSubarraySum([]int{23, 2, 6, 4, 7}, 13)
	if got != want {
		t.Errorf("test checkSubarraySum error, got: %v, want: %v", got, want)
	}

	want = false
	got = checkSubarraySum([]int{1, 2, 12}, 6)
	if got != want {
		t.Errorf("test checkSubarraySum error, got: %v, want: %v", got, want)
	}

	want = false
	got = checkSubarraySum([]int{0}, 6)
	if got != want {
		t.Errorf("test checkSubarraySum error, got: %v, want: %v", got, want)
	}

	want = true
	got = checkSubarraySum([]int{5,0,0,0}, 6)
	if got != want {
		t.Errorf("test checkSubarraySum error, got: %v, want: %v", got, want)
	}
}
