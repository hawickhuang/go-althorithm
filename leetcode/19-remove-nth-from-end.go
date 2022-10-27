package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head

	i := 0
	for fast != nil {
		fast = fast.Next
		if i >= n {
			slow = slow.Next
		}
		i++
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
