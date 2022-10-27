package leetcode

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			inode1 := head
			inode2 := fast
			for inode1 != inode2 {
				inode1 = inode1.Next
				inode2 = inode2.Next
			}
			return inode1
		}
	}
	return nil
}
