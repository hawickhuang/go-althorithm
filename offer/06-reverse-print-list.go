package offer

type ListNode struct {
	Next *ListNode
	Val int
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var res []int
	var helper func(node *ListNode)
	helper = func(node *ListNode) {
		if node == nil {
			return
		}
		helper(node.Next)
		res = append(res, node.Val)
	}
	helper(head)

	return res
}
