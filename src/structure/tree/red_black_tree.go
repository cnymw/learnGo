package tree

type RedBlackTree struct {
	Root *Node
}

func (tree *RedBlackTree) leftRotate(x *Node) {
	// 设置y
	y := x.Right

	// 双向绑定x与y左子树:y.Left
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}

	// 双向绑定y与x父节点:x.Parent
	y.Parent = x.Parent
	if x.Parent == nil {
		tree.Root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	// 目前y.Left仍然为老的左子树，需要重新与x进行绑定，双向绑定x与y左子树:y.Left
	y.Left = x
	x.Parent = y
}
