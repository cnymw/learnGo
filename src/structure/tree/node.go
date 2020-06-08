package tree

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewEmptyNode(v int) *Node {
	return &Node{Key: v}
}
