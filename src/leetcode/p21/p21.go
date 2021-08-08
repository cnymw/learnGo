package p21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2 := l1, l2
	result := &ListNode{}
	cur := result
	for cur1 != nil && cur2 != nil {
		if cur1.Val >= cur2.Val {
			cur.Next = cur2
			cur = cur.Next
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur = cur.Next
			cur1 = cur1.Next
		}
	}
	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return result.Next
}
