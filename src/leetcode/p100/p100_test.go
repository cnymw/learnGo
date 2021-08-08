package p100

import (
	"fmt"
	"testing"
)

//  [1,2,3],   [1,2,3]
func TestIsSameTree(t *testing.T) {
	node13 := TreeNode{3, nil, nil}
	node12 := TreeNode{2, nil, nil}
	node11 := TreeNode{1, &node12, &node13}

	node23 := TreeNode{3, nil, nil}
	node22 := TreeNode{2, nil, nil}
	node21 := TreeNode{1, &node22, &node23}

	fmt.Println(isSameTree(&node11, &node21))
}

//  [1,2],     [1,null,2]
func TestIsSameTree2(t *testing.T) {
	node12 := TreeNode{2, nil, nil}
	node11 := TreeNode{1, &node12, nil}

	node22 := TreeNode{2, nil, nil}
	node21 := TreeNode{1, nil, &node22}

	fmt.Println(isSameTree(&node11, &node21))
}

//  [1,2,1],   [1,1,2]
func TestIsSameTree3(t *testing.T) {
	node13 := TreeNode{1, nil, nil}
	node12 := TreeNode{2, nil, nil}
	node11 := TreeNode{1, &node12, &node13}

	node23 := TreeNode{2, nil, nil}
	node22 := TreeNode{1, nil, nil}
	node21 := TreeNode{1, &node22, &node23}

	fmt.Println(isSameTree(&node11, &node21))
}
