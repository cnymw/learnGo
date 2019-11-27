package tree

import (
	"fmt"
	"testing"
)

func TestSearchTree_Insert(t *testing.T) {
	tree := &SearchTree{}
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(13)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(15)
	tree.PrintTree()
}

func TestSearchTree_Contains(t *testing.T) {
	tree := &SearchTree{}
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

func TestSearchTree_FindMin(t *testing.T) {
	tree := &SearchTree{}
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(13)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(15)
	fmt.Println(tree.FindMax().Val)
	fmt.Println(tree.FindMin().Val)
}

func TestSearchTree_Remove(t *testing.T) {
	tree := &SearchTree{}
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
