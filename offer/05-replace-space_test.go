package offer

import "testing"

func Test_replaceSpace(t *testing.T) {
	s := replaceSpace("s a b c d")
	if s != "s%20a%20b%20c%20d" {
		t.Errorf("test replaceSpace() error, got: %s, want: %s", s, "s%20a%20b%20c%20d")
	}
	s = replaceSpace("  d ")
	if s != "%20%20d%20" {
		t.Errorf("test replaceSpace() error, got: %s, want: %s", s, "%20%20d%20")
	}
	s = replaceSpace("  ")
	if s != "%20%20" {
		t.Errorf("test replaceSpace() error, got: %s, want: %s", s, "%20%20")
	}
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val: nums[0],
		Next: buildList(nums[1:]),
	}
	return head
}

func compareList(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	l := len(nums2)
	for i := range nums1{
		if nums1[i] != nums2[l-i-1] {
			return false
		}
	}
	return true
}

func Test_reversePrint(t *testing.T) {
	t1 := []int{1, 2, 3}
	head := buildList(t1)
	r := reversePrint(head)
	if !compareList(t1, r) {
		t.Errorf("test error for reversePrint(), got: %v", r )
	}
}
