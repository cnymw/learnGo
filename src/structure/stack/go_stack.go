package stack

import "errors"

type Stack struct {
	S   []int
	Top int64
}

func NewStack(size int) *Stack {
	return &Stack{
		S:   make([]int, size),
		Top: -1,
	}
}
func (stack *Stack) IsEmpty() bool {
	if stack.Top == -1 {
		return true
	}
	return false
}

func (stack *Stack) Push(v int) {
	stack.Top++
	stack.S[stack.Top] = v
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("underflow")
	}
	stack.Top--
	return stack.S[stack.Top+1], nil
}
