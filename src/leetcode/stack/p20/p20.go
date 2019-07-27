package p20

import "errors"

type Stack struct {
	Val []rune
	Top int64
}

func NewStack(size int) *Stack {
	return &Stack{
		Val: make([]rune, size),
		Top: -1,
	}
}

func (stack *Stack) Push(v rune) {
	stack.Top++
	stack.Val[stack.Top] = v
}

func (stack *Stack) IsEmpty() bool {
	if stack.Top == -1 {
		return true
	}
	return false
}

func (stack *Stack) Pop() (rune, error) {
	if stack.IsEmpty() {
		return 0, errors.New("underflow")
	}
	stack.Top--
	return stack.Val[stack.Top+1], nil
}

func isValid(s string) bool {
	array := []rune(s)
	if len(array)%2 != 0 {
		return false
	}
	stack := NewStack(1000000)
	for index := 0; index < len(array); index++ {
		r := array[index]
		if isLeft(r) {
			stack.Push(r)
		} else if isRight(r) {
			pop, err := stack.Pop()
			if err != nil {
				return false
			}
			if !isMatch(pop, r) {
				return false
			}
		} else {
			return false
		}
	}
	if stack.IsEmpty() {
		return true
	} else {
		return false
	}

}
func isLeft(r rune) bool {
	if r == '(' || r == '[' || r == '{' {
		return true
	}
	return false
}

func isRight(r rune) bool {
	if r == ')' || r == ']' || r == '}' {
		return true
	}
	return false
}

func isMatch(l, r rune) bool {
	if l == '(' && r == ')' {
		return true
	} else if l == '[' && r == ']' {
		return true
	} else if l == '{' && r == '}' {
		return true
	} else {
		return false
	}
}
