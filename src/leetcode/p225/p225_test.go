package p225

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	k := Constructor()
	k.Push(1)
	k.Push(2)

	v := k.Top()
	fmt.Println(v)

	v = k.Pop()
	fmt.Println(v)

	fmt.Println(k.Empty())
}
