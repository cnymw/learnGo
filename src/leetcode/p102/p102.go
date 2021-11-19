package p102

import (
	"learnGo/src/leetcode/tool"
)

func levelOrder(root *tool.TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*tool.TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})

		p := []*tool.TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}

	return ret
}
