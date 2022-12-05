package p92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 创建临时数组，用于链表数据同步到数组
	array := make([]int, 0)
	cur := head

	// 通过 cur 遍历链表，将链表数据同步到数组
	for cur != nil {
		array = append(array, cur.Val)
		cur = cur.Next
	}

	// 利用数组特性：随机读取，通过该特性可以将指定的两个数组元素进行调换
	for i := 0; i < (right-left+1)/2; i++ {
		array[left+i-1], array[right-i-1] = array[right-i-1], array[left+i-1]
	}

	cur = head
	// 通过 cur 遍历链表，将数组元素同步到链表
	for i := 0; i < len(array); i++ {
		cur.Val = array[i]
		cur = cur.Next
	}
	return head
}
