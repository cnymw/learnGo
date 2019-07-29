package p232

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)

	v := queue.Peek()
	fmt.Println(v)

	v = queue.Pop()
	fmt.Println(v)

	v = queue.Pop()
	fmt.Println(v)

	fmt.Println(queue.Empty())
}

func TestNew1(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	v := queue.Pop()
	fmt.Println(v)

	v = queue.Pop()
	fmt.Println(v)

	v = queue.Pop()
	fmt.Println(v)

	fmt.Println(queue.Empty())
}
