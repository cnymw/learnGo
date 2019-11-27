package tree

import "fmt"

// 方法命名，以及接口实现基于算法导论

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree interface {
	inorderTreeWalk(node *Node)
	treeSearch(node *Node, k int) *Node
}

type BinaryTree struct{}

func (entry *BinaryTree) inorderTreeWalk(node *Node) {
	if node != nil {
		entry.inorderTreeWalk(node.Left)
		fmt.Println(node.Val)
		entry.inorderTreeWalk(node.Right)
	}
}

func (entry *BinaryTree) treeSearch(node *Node, k int) *Node {
	if node == nil || k == node.Val {
		return node
	}
	if k < node.Val {
		return entry.treeSearch(node.Left, k)
	} else {
		return entry.treeSearch(node.Right, k)
	}
}
