package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(13)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(15)
	tree.PrintTree()
}

func TestBinarySearchTree_Contains(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(13)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(15)
	fmt.Println(tree.Contains(13))
	fmt.Println(tree.Contains(22))
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(13)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(15)
	fmt.Println(tree.FindMax().Key)
	fmt.Println(tree.FindMin().Key)
}

func TestBinarySearchTree_Remove(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)
	tree.PrintTree()
	fmt.Println("after remove")
	tree.Remove(2)
	tree.PrintTree()
}

func TestBinarySearchTree_TreeSuccessor(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)

	x := tree.TreeSearch(5)
	fmt.Println(tree.TreeSuccessor(x).Key)
}
