package p92

import (
	"fmt"
	"testing"
)

func TestReverseBetween1(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	node := reverseBetween(node1, 2, 4)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestReverseBetween2(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node3 := &ListNode{Val: 3, Next: node5}

	node := reverseBetween(node3, 1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
