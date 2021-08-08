package p145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	travel := make([]int, 0)
	if root.Left != nil {
		travel = append(travel, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		travel = append(travel, postorderTraversal(root.Right)...)
	}
	travel = append(travel, root.Val)
	return travel
}
