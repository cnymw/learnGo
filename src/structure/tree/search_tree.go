package tree

import "fmt"

func NewEmptyNode(v int) *Node {
	return &Node{Val: v}
}

type Search interface {
	MakeEmpty()
	IsEmpty() bool
	Contains(v int) bool
	FindMin() *Node
	FindMax() *Node
	Insert(v int)
	Remove(v int)
	PrintTree()
}

type SearchTree struct {
	root *Node
}

func (tree *SearchTree) MakeEmpty() {
	tree.root = nil
}
func (tree *SearchTree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *SearchTree) Contains(v int) bool {
	return contains(v, tree.root)
}

func contains(v int, node *Node) bool {
	if node == nil {
		return false
	}
	if v > node.Val {
		return contains(v, node.Right)
	} else if v < node.Val {
		return contains(v, node.Left)
	} else {
		return true
	}
}

func (tree *SearchTree) FindMin() *Node {
	if tree.IsEmpty() {
		return nil
	}
	return findMin(tree.root)
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (tree *SearchTree) FindMax() *Node {
	if tree.IsEmpty() {
		return nil
	}
	return findMax(tree.root)
}
func findMax(node *Node) *Node {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *SearchTree) Insert(v int) {
	tree.root = insert(v, tree.root)
}

func insert(v int, node *Node) *Node {
	if node == nil {
		return NewEmptyNode(v)
	}
	if v > node.Val {
		node.Right = insert(v, node.Right)
	} else if v < node.Val {
		node.Left = insert(v, node.Left)
	} else {

	}
	return node
}

func (tree *SearchTree) PrintTree() {
	if tree.IsEmpty() {
		fmt.Println("empty tree")
	} else {
		printTree(tree.root)
	}
}

func printTree(node *Node) {
	if node != nil {
		printTree(node.Left)
		fmt.Println(node.Val)
		printTree(node.Right)
	}
}

func (tree *SearchTree) Remove(v int) {
	tree.root = remove(v, tree.root)
}

func remove(v int, node *Node) *Node {
	if node == nil {
		return node
	}
	if v > node.Val {
		node.Right = remove(v, node.Right)
	} else if v < node.Val {
		node.Left = remove(v, node.Left)
	} else if node.Left != nil && node.Right != nil {
		node.Val = findMin(node.Right).Val
		node.Right = remove(node.Val, node.Right)
	} else {
		if node.Left == nil {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}
