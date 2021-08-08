package p2

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	node14 := &ListNode{Val: 3}
	node12 := &ListNode{Val: 4, Next: node14}
	node11 := &ListNode{Val: 2, Next: node12}

	node24 := &ListNode{Val: 4}
	node23 := &ListNode{Val: 6, Next: node24}
	node21 := &ListNode{Val: 5, Next: node23}

	node := addTwoNumbers(node11, node21)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
