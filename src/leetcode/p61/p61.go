package p61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	tail, count := head, 1

	// 遍历链表以指向尾部 tail
	for tail.Next != nil {
		tail = tail.Next
		count++
	}

	// 将尾部 tail 的 Next 指向头部 head，以实现单向循环链表结构
	tail.Next = head

	// 防止移动数大于链表节点数量
	k = k % count

	k = count - 1 - k
	node := head

	// 找到新的头部节点的前一个节点
	for i := 0; i < k; i++ {
		node = node.Next
	}
	// 设置链表的新的头部节点
	head = node.Next
	// 将链表新的尾部节点的 Next 置为空，将单向循环链表重新设置为单向链表
	node.Next = nil
	return head
}
