package list

import (
	"fmt"
	"testing"
)

func TestDoubleList_Display(t *testing.T) {
	list := DoubleList{}
	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})
	list.Append(&DoubleNode{Data: 3})
	list.Append(&DoubleNode{Data: 4})

	list.Display()
}

func TestDoubleList_Get(t *testing.T) {
	list := DoubleList{}
	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})
	list.Append(&DoubleNode{Data: 3})
	list.Append(&DoubleNode{Data: 4})

	fmt.Println(list.Get(2))
}

func TestDoubleList_Insert(t *testing.T) {
	list := DoubleList{}
	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})
	list.Append(&DoubleNode{Data: 3})
	list.Append(&DoubleNode{Data: 4})
	list.Insert(2, &DoubleNode{Data: 5})

	list.Display()
}

func TestDoubleList_Delete(t *testing.T) {
	list := DoubleList{}
	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})
	list.Append(&DoubleNode{Data: 3})
	list.Append(&DoubleNode{Data: 4})

	list.Delete(2)
	list.Display()
}

func TestDoubleList_Reverse(t *testing.T) {
	list := DoubleList{}
	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})
	list.Append(&DoubleNode{Data: 3})
	list.Append(&DoubleNode{Data: 4})

	list.Reverse()
}
