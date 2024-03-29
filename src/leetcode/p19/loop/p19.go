package loop

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	cur := head
	// 遍历 ListNode，获得链表长度 length
	for cur != nil {
		length++
		cur = cur.Next
	}
	dummy := &ListNode{Val: 0, Next: head}
	cur = dummy
	// 遍历到待删除节点的前一个节点
	for i := 1; i < length-n+1; i++ {
		cur = cur.Next
	}
	// 将待删除的节点的前一个节点的 Next 指向待删除的节点的 Next
	// 这样就将待删除节点给删除了
	// 例如 a->b-c 链表，
	// cur 为 a，
	// 将 a.Next 指向 c，
	// 也就是 cur.Next = b.next = cur.Next.Next
	cur.Next = cur.Next.Next
	return dummy.Next
}
