package p100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	pstack, qstack := New(), New()

	if p != nil {
		pstack = pstack.Push(p)
	}
	if q != nil {
		qstack = qstack.Push(q)
	}

	var ppop, qpop interface{}
	for !pstack.Empty() && !qstack.Empty() {
		pstack, ppop = pstack.Pop()
		qstack, qpop = qstack.Pop()

		ptemp, qtemp := ppop.(*TreeNode), qpop.(*TreeNode)
		if ptemp.Val != qtemp.Val {
			return false
		}

		if (ptemp.Left == nil && qtemp.Left != nil) || (ptemp.Left != nil && qtemp.Left == nil) {
			return false
		} else {
			if ptemp.Left != nil && qtemp.Left != nil {
				pstack = pstack.Push(ptemp.Left)
				qstack = qstack.Push(qtemp.Left)
			}
		}

		if (ptemp.Right == nil && qtemp.Right != nil) || (ptemp.Right != nil && qtemp.Right == nil) {
			return false
		} else {
			if ptemp.Right != nil && qtemp.Right != nil {
				pstack = pstack.Push(ptemp.Right)
				qstack = qstack.Push(qtemp.Right)
			}
		}
	}

	if !pstack.Empty() || !qstack.Empty() {
		return false
	}
	return true
}

type stack []interface{}

func New() stack {
	s := make(stack, 0)
	return s
}
func (s stack) Push(v interface{}) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, interface{}) {
	if s.Len() == 0 {
		return s, nil
	}
	return s[:s.Len()-1], s[s.Len()-1]
}

func (s stack) Len() int {
	return len(s)
}

func (s stack) Peek() interface{} {
	return s[s.Len()-1]
}

func (s stack) Empty() bool {
	return s.Len() == 0
}
