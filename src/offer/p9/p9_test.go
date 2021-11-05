package p9

import (
	"fmt"
	"testing"
)

func TestP9_0(t *testing.T) {
	cqueue := &CQueue{}
	cqueue.AppendTail(3)
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
}

func TestP9_1(t *testing.T) {
	cqueue := &CQueue{}
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(5)
	cqueue.AppendTail(2)
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
}

func TestP9_2(t *testing.T) {
	cqueue := &CQueue{}
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(12)
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(10)
	cqueue.AppendTail(9)
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(20)
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(1)
	cqueue.AppendTail(8)
	cqueue.AppendTail(20)
	cqueue.AppendTail(1)
	cqueue.AppendTail(11)
	cqueue.AppendTail(2)
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
}

func TestP9_3(t *testing.T) {
	cqueue := &CQueue{}
	cqueue.AppendTail(3)
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
}
