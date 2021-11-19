package p111

import (
	"fmt"
	"learnGo/src/leetcode/tool"
	"testing"
)

func TestMinDepth(t *testing.T) {
	array := []int{3, 9, 20, 0, 0, 15, 7}
	root := tool.ArrayToTree(array)
	fmt.Println(minDepth(root))
}
