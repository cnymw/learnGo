package p94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	travel := make([]int, 0)
	if root.Left != nil {
		travel = append(travel, inorderTraversal(root.Left)...)
	}
	travel = append(travel, root.Val)
	if root.Right != nil {
		travel = append(travel, inorderTraversal(root.Right)...)
	}
	return travel
}
