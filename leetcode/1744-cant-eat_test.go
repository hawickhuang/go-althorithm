package leetcode

import (
	"testing"
)

func Test_canEat(t *testing.T) {
	candies := []int{7, 4, 5, 3, 8}
	queries := [][]int{{0, 2, 2}, {4, 2, 4}, {2, 3, 1000000000}}
	r := canEat(candies, queries)
	w := []bool{true, false, true}
	if !comapre(r, w) {
		t.Errorf("test canEat() error: got: %v, want: %v", r, w)
	}
	candies = []int{5, 2, 6, 4, 1}
	queries = [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}}
	r = canEat(candies, queries)
	w = []bool{false, true, true, false, false}
	if !comapre(r, w) {
		t.Errorf("test canEat() error: got: %v, want: %v", r, w)
	}
}

func comapre(b1, b2 []bool) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
