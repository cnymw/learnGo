package p1038

import (
	"fmt"
	"testing"
)

func TestBstToGst(t *testing.T) {
	node0 := &TreeNode{Val: 0}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Right: node3}
	node1 := &TreeNode{Val: 1, Left: node0, Right: node2}

	node5 := &TreeNode{Val: 5}
	node8 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7, Right: node8}
	node6 := &TreeNode{Val: 6, Left: node5, Right: node7}

	node4 := &TreeNode{Val: 4, Left: node1, Right: node6}

	bstToGst(node4)
	inorderTreeWalk_Test(node4)
}

func inorderTreeWalk_Test(node *TreeNode) {
	if node != nil {
		inorderTreeWalk_Test(node.Left)
		fmt.Println(node.Val)
		inorderTreeWalk_Test(node.Right)
	}
}
