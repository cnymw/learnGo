package sdk_list

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	node := removeNthFromEnd(node1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
