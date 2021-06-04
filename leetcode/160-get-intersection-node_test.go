package leetcode

import (
	"testing"
)

type testCase160 struct {
	headA *ListNode
	headB *ListNode
	want  *ListNode
}

func Test_getIntersectionNode(t *testing.T) {
	l1, l11 := BuildList(8, []int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}, 2, 3)
	l2, l22 := BuildList(2, []int{0, 9, 1, 2, 4}, []int{3, 2, 4}, 3, 1)
	l3, l33 := BuildList(0, []int{2,6,4}, []int{1,5}, 3, 2)
	cases := []testCase160{
		{headA: l1, headB: l11, want: &ListNode{Val: 8, Next: nil}},
		{headA: l2, headB: l22, want: &ListNode{Val: 2, Next: nil}},
		{headA: l2, headB: l22, want: &ListNode{Val: 2, Next: nil}},
		{headA: l3, headB: l33, want: nil},
	}
	for _, c := range cases {
		got := getIntersectionNode(c.headA, c.headB)
		if (got == nil && c.want != nil ) || (got != nil && got.Val != c.want.Val) {
			t.Errorf("test error for input %v, got: %v, want: %v", c, got, c.want)
		}
	}
}

func BuildList(intersectVal int, listA, listB []int, skipA, skipB int) (*ListNode, *ListNode) {
	la := len(listA)
	lb := len(listB)

	dummyA := &ListNode{Val: -1}
	dummyB := &ListNode{Val: -1}

	headA := dummyA
	headB := dummyB

	for i := 0; i < la; i++ {
		if i == skipA {
			break
		}
		next := &ListNode{Val: listA[i]}
		headA.Next = next
		headA = next
	}
	for i := 0; i < lb; i++ {
		if i == skipB {
			break
		}
		next := &ListNode{Val: listB[i]}
		headB.Next = next
		headB = next
	}

	for i := skipA; i < la; i++ {
		next := &ListNode{Val: listA[i]}
		if i == skipA {
			headB.Next = next
		}
		headA.Next = next
		headA = next
	}

	return dummyA.Next, dummyB.Next
}
