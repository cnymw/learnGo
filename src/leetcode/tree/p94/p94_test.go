package p94

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	node3 := TreeNode{3, nil, nil}
	node2 := TreeNode{2, &node3, nil}
	node1 := TreeNode{1, nil, &node2}

	traval := inorderTraversal(&node1)
	for _, v := range traval {
		fmt.Println(v)
	}
}
