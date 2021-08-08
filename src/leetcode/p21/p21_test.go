package p21

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	node14 := &ListNode{Val: 4}
	node12 := &ListNode{Val: 2, Next: node14}
	node11 := &ListNode{Val: 1, Next: node12}

	node24 := &ListNode{Val: 4}
	node23 := &ListNode{Val: 3, Next: node24}
	node21 := &ListNode{Val: 1, Next: node23}

	node := mergeTwoLists(node11, node21)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
