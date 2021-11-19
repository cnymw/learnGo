package p104

import (
	"fmt"
	"learnGo/src/leetcode/tool"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	array := []int{3, 9, 20, 0, 0, 15, 7}
	root := tool.ArrayToTree(array)
	fmt.Println(maxDepth(root))
}
