package tree

import (
	"fmt"
)

type Tree interface {
	InorderTreeWalk(*Node)
	PreorderTreeWalk(*Node)
	PostorderTreeWalk(*Node)
	TreeSearch(int) *Node
	TreeSuccessor(*Node) *Node
	TreePredecessor(*Node) *Node
	MakeEmpty()
	IsEmpty() bool
	Contains(int) bool
	FindMin() *Node
	FindMax() *Node
	Insert(int)
	PrintTree()
	Remove(int)
}

type BinarySearchTree struct {
	Root *Node
}

func (t *BinarySearchTree) TreeSuccessor(x *Node) *Node {
	if x == nil {
		return nil
	}
	if x.Right != nil {
		return findMin(x.Right)
	} else {
		y := x.Parent
		for y != nil && x == y.Right {
			x = y
			y = y.Parent
		}
		return y
	}
}

func (t *BinarySearchTree) TreePredecessor(x *Node) *Node {
	if x == nil {
		return nil
	}
	if x.Left != nil {
		return findMax(x.Left)
	} else {
		y := x.Parent
		for y != nil && x == y.Left {
			x = y
			y = y.Parent
		}
		return y
	}
}

func (t *BinarySearchTree) MakeEmpty() {
	t.Root = nil
}
func (t *BinarySearchTree) IsEmpty() bool {
	return t.Root == nil
}

func (t *BinarySearchTree) Contains(v int) bool {
	return contains(v, t.Root)
}

func contains(v int, x *Node) bool {
	if x == nil {
		return false
	}
	if v > x.Key {
		return contains(v, x.Right)
	} else if v < x.Key {
		return contains(v, x.Left)
	} else {
		return true
	}
}

func (t *BinarySearchTree) FindMin() *Node {
	if t.IsEmpty() {
		return nil
	}
	return findMin(t.Root)
}

func findMin(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func (t *BinarySearchTree) FindMax() *Node {
	if t.IsEmpty() {
		return nil
	}
	return findMax(t.Root)
}
func findMax(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

func (t *BinarySearchTree) Insert(v int) {
	z := NewEmptyNode(v)
	x := t.Root
	var y *Node

	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.Parent = y
	if y == nil {
		t.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}

func (t *BinarySearchTree) PrintTree() {
	if t.IsEmpty() {
		fmt.Println("empty t")
	} else {
		printTree(t.Root)
	}
}

func printTree(x *Node) {
	if x != nil {
		printTree(x.Left)
		fmt.Println(x.Key)
		printTree(x.Right)
	}
}

func (t *BinarySearchTree) Remove(v int) {
	t.Root = remove(v, t.Root)
}

func remove(v int, x *Node) *Node {
	if x == nil {
		return x
	}
	if v > x.Key {
		x.Right = remove(v, x.Right)
	} else if v < x.Key {
		x.Left = remove(v, x.Left)
	} else if x.Left != nil && x.Right != nil {
		x.Key = findMin(x.Right).Key
		x.Right = remove(x.Key, x.Right)
	} else {
		if x.Left == nil {
			x = x.Right
		} else {
			x = x.Left
		}
	}
	return x
}

func (t *BinarySearchTree) InorderTreeWalk(x *Node) {
	if x != nil {
		t.InorderTreeWalk(x.Left)
		fmt.Println(x.Key)
		t.InorderTreeWalk(x.Right)
	}
}

func (t *BinarySearchTree) PreorderTreeWalk(x *Node) {
	if x != nil {
		fmt.Println(x.Key)
		t.PreorderTreeWalk(x.Left)
		t.PreorderTreeWalk(x.Right)
	}
}

func (t *BinarySearchTree) PostorderTreeWalk(x *Node) {
	if x != nil {
		t.PostorderTreeWalk(x.Left)
		t.PostorderTreeWalk(x.Right)
		fmt.Println(x.Key)
	}
}

func (t *BinarySearchTree) TreeSearch(k int) *Node {
	return treeSearch(t.Root, k)
}

func treeSearch(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return treeSearch(x.Left, k)
	} else {
		return treeSearch(x.Right, k)
	}
}
