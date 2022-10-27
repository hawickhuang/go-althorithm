package leetcode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for {
		if cur== nil {
			break
		}
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
