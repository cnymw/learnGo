package p1038

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	inorderTreeWalk(root, &sum)
	return root
}

func inorderTreeWalk(node *TreeNode, sum *int) {
	if node != nil {
		inorderTreeWalk(node.Right, sum)
		if *sum != 0 {
			node.Val += *sum
		}
		*sum = node.Val
		inorderTreeWalk(node.Left, sum)
	}
}
