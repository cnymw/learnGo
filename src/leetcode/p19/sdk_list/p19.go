package sdk_list

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 通过使用 Golang SDK 新建一个双向循环链表
	l := list.New()
	// 将 ListNode 结构复制到链表里面
	for head != nil {
		l.PushBack(head.Val)
		head = head.Next
	}
	remove := l.Back()

	// 从双向循环链表的末尾向前数 n 个节点就是想要删除的节点
	for i := 1; i < n; i++ {
		remove = remove.Prev()
	}
	// 调用 list SDK 接口实现删除逻辑
	l.Remove(remove)

	// 将双向循环链表数据重新复制到 ListNode 结构
	if l.Len() > 0 {
		var newHead, tail *ListNode
		element := l.Front()
		newHead = &ListNode{Val: element.Value.(int)}
		tail = newHead
		// 从双向循环链表的首节点开始遍历
		for i := 1; i < l.Len(); i++ {
			element = element.Next()
			tail.Next = &ListNode{Val: element.Value.(int)}
			tail = tail.Next
		}
		return newHead
	} else {
		return nil
	}
}
