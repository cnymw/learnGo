package tree

type RedBlackTree struct {
	Root *Node
}

func (tree *RedBlackTree) leftRotate(x *Node) {
	// 设置 y 为 x 的右子树
	y := x.Right

	// 双向绑定 x 与 y 左子树 y.Left
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}

	// 双向绑定 y 与 x 父节点 x.Parent
	y.Parent = x.Parent
	if x.Parent == nil {
		tree.Root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	// 目前 y.Left 仍然为老的左子树，需要重新与 x 进行绑定，双向绑定 x 与 y 左子树 y.Left
	y.Left = x
	x.Parent = y
}

func (tree *RedBlackTree) rightRotate(x *Node) {
	// 设置 y 为 x 的左子树
	y := x.Left

	// 双向绑定 x 与 y 右子树 y.Right
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}

	// 双向绑定 y 与 x 父节点 x.Parent
	y.Parent = x.Parent
	if x.Parent == nil {
		tree.Root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	// 目前 y.Right 仍然为老的右子树，需要重新与 x 进行绑定，双向绑定 x 与 y 右子树 y.Right
	y.Right = x
	x.Parent = y
}
