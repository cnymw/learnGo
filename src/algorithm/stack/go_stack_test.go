package stack

import (
	"fmt"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack(10)
	fmt.Println(stack.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	stack := NewStack(10)
	stack.Push(10)
	stack.Push(11)
	stack.Push(12)
	for _, v := range stack.S {
		fmt.Println(v)
	}
	fmt.Println(stack.IsEmpty())
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack(10)
	stack.Push(10)
	stack.Push(11)
	stack.Push(12)
	for _, v := range stack.S {
		fmt.Println(v)
	}
	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
}
