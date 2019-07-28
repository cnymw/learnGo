package p144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	travel := make([]int, 0)
	travel = append(travel, root.Val)
	if root.Left != nil {
		travel = append(travel, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		travel = append(travel, preorderTraversal(root.Right)...)
	}
	return travel
}
