package p101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	} else {
		return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
	}
}
