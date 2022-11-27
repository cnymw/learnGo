package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// head 链表首节点
	// tail 链表尾节点
	var head, tail *ListNode

	// 记录进位值
	var carry int

	// 持续遍历链表 l1,l2，直到两个链表都遍历完毕
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		// 如果结果链表为空，那么创建链表首节点
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 如果进位值不为 0，那么在链表尾节点新增一个节点，用于记录进位值
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	// 返回链表首节点
	return head
}
