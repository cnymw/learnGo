package tree

import (
	"fmt"
	"testing"
)

func TestInorderTreeWalk(t *testing.T) {
	entry := &BinaryTree{}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	node3 := &Node{Val: 3, Left: node2, Right: node4}

	node6 := &Node{Val: 6}
	node8 := &Node{Val: 8}
	node7 := &Node{Val: 7, Left: node6, Right: node8}

	node5 := &Node{Val: 5, Left: node3, Right: node7}
	entry.inorderTreeWalk(node5)
}

func TestTreeSearch(t *testing.T) {
	entry := &BinaryTree{}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	node3 := &Node{Val: 3, Left: node2, Right: node4}

	node6 := &Node{Val: 6}
	node8 := &Node{Val: 8}
	node7 := &Node{Val: 7, Left: node6, Right: node8}

	node5 := &Node{Val: 5, Left: node3, Right: node7}
	fmt.Println(entry.treeSearch(node5, 3))
}
